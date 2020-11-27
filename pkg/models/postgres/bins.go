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

// Destroy a bin - delete it from the database
func (m *BinModel) Destroy(binID string) (string, error) {
	stmt := `DELETE FROM bins
					 WHERE id = $1
					 RETURNING id`

	var deletedID string
	err := m.DB.QueryRow(stmt, binID).Scan(&deletedID)
	if err != nil {
		return "", err
	}

	return deletedID, nil
}

// GetUserBins returns a slice of all of the bins for a specified UserID
func (m *BinModel) GetUserBins(userID int) ([]*models.Bin, error) {
	stmt := `SELECT id, user_id, created
					 FROM bins
					 WHERE user_id = $1`

	rows, err := m.DB.Query(stmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bins := []*models.Bin{}
	for rows.Next() {
		b := &models.Bin{}
		err = rows.Scan(&b.ID, &b.UserID, &b.Created)
		if err != nil {
			return nil, err
		}

		bins = append(bins, b)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return bins, nil
}
