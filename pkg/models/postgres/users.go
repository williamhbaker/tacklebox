package postgres

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/lib/pq"
	"github.com/wbaker85/tacklebox/pkg/models"
)

// UserModel wraps the database
type UserModel struct {
	DB *sql.DB
}

// Insert a new user into the database
func (m *UserModel) Insert(email, password string) error {
	stmt := `INSERT INTO users (email, hashed_password, created)
					 VALUES ($1, $2, CURRENT_TIMESTAMP)`

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	_, err = m.DB.Exec(stmt, email, hash)
	if err != nil {
		var pgError *pq.Error
		if errors.As(err, &pgError) {
			if pgError.Code.Name() == "unique_violation" {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

// Authenticate an existing user per the supplied email and password
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get a user representation from an ID
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
