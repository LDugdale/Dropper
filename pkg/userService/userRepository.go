package userService

import (
    "github.com/ldugdale/dropper/pkg/database"
    "github.com/LDugdale/Dropper/pkg/types"
    "database/sql"
)

type Book struct {
    Isbn   string
    Title  string
    Author string
    Price  float32
}

type UserRepository struct {
	db database.DB
}

func NewUserRepository(db database.DB) *UserRepository {
    return &UserRepository{
        db: db,
    }
}

func (ur *UserRepository) CreateUser(user *types.UserModel) (int64, error) {
    
    insertUserStatement := `
        INSERT INTO Users (username, password)
        SELECT '$1', '$2'
        WHERE NOT EXISTS (
            SELECT 
                username,
                password
            FROM Users
            WHERE username = ($1)
            LIMIT 1
        );
    `

    result, err := ur.db.Exec(insertUserStatement, user.Username, user.Password)
    if err != nil {
      return -1, err
    }
  
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return -1, err
    }
  
    return rowsAffected, nil
}

func (ur *UserRepository) GetUser(username string) (*types.User, error) {

    row := ur.db.QueryRow("SELECT * FROM Users WHERE Username = $1 LIMIT 1", username)

    user := new(types.User)
    err := row.Scan(&user.Username)
    if err == sql.ErrNoRows {
      return nil, err
    } else if err != nil {
      return nil, err
    }

    return user, nil
}

