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
func (m *UserModel) Insert(email, password string) (int, error) {
	stmt := `INSERT INTO users (email, hashed_password, created)
					 VALUES ($1, $2, CURRENT_TIMESTAMP)
					 RETURNING id`

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, err
	}

	var id int
	err = m.DB.QueryRow(stmt, email, hash).Scan(&id)
	if err != nil {
		var pgError *pq.Error
		if errors.As(err, &pgError) {
			if pgError.Code.Name() == "unique_violation" {
				return 0, models.ErrDuplicateEmail
			}
		}
		return 0, err
	}

	return id, nil
}

// Authenticate an existing user per the supplied email and password
func (m *UserModel) Authenticate(email, password string) (int, error) {
	stmt := `SELECT id, hashed_password FROM users WHERE email = $1`

	var id int
	var hashedPassword []byte

	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}

	return id, nil
}

// Get a user representation from an ID
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
