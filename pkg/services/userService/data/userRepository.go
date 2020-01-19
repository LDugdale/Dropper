package data

import (
	"database/sql"

	"github.com/ldugdale/dropper/pkg/commonAbstractions"
	"github.com/ldugdale/dropper/pkg/database"
	"github.com/ldugdale/dropper/pkg/log"
)

type UserRepository struct {
	logger log.Logger
	db     database.DB
}

func NewUserRepository(logger log.Logger, db database.DB) *UserRepository {
	return &UserRepository{
		logger: logger,
		db:     db,
	}
}

func (ur *UserRepository) CreateUser(user *commonAbstractions.UserModel) (int64, error) {

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

func (ur *UserRepository) GetUser(username string) (*commonAbstractions.UserModel, error) {

	row := ur.db.QueryRow("SELECT * FROM Users WHERE Username = ? LIMIT 1", username)

	user := new(commonAbstractions.UserModel)
	err := row.Scan(&user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
