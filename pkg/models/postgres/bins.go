package postgres

import (
	"database/sql"
	"errors"

	"github.com/wbaker85/tacklebox/pkg/models"

	"github.com/lib/pq"
)

// BinModel wraps the database used for the bin methods
type BinModel struct {
	DB *sql.DB
}

// Insert a new bin into the database
func (m *BinModel) Insert(id string, userID int) (string, error) {
	stmt := `INSERT INTO bins (id, user_id, created)
					 VALUES ($1, $2, CURRENT_TIMESTAMP)
					 RETURNING id`

	var createdID string
	err := m.DB.QueryRow(stmt, id, userID).Scan(&createdID)
	if err != nil {
		var pgError *pq.Error
		if errors.As(err, &pgError) {
			if pgError.Code.Name() == "foreign_key_violation" {
				return "", models.ErrInvalidUser
			}
		}

		return "", err
	}

	return createdID, nil
}
