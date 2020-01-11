package data

import (
    "github.com/ldugdale/dropper/pkg/database"
    "github.com/LDugdale/Dropper/pkg/services/userService/abstractions"
    "database/sql"
)

type UserRepository struct {
	db database.DB
}

func NewUserRepository(db database.DB) *UserRepository {
    return &UserRepository{
        db: db,
    }
}

func (ur *UserRepository) CreateUser(user *abstractions.UserModel) (int64, error) {

    insertUserStatement := `
        INSERT INTO Users (username, password)
        SELECT ? , ?
        WHERE NOT EXISTS (
            SELECT 
                username,
                password
            FROM Users
            WHERE username = ?
            LIMIT 1
        );
    `

    result, err := ur.db.Exec(insertUserStatement, user.Username, user.Password, user.Username)
    if err != nil {
      return -1, err
    }
  
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return -1, err
    }
  
    return rowsAffected, nil
}

func (ur *UserRepository) GetUser(username string) (*abstractions.UserModel, error) {

    row := ur.db.QueryRow("SELECT * FROM Users WHERE Username = $1 LIMIT 1", username)

    user := new(abstractions.UserModel)
    err := row.Scan(&user.Username, &user.Password)
    if err == sql.ErrNoRows {
      return nil, err
    } else if err != nil {
      return nil, err
    }

    return user, nil
}

