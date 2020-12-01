package postgres

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/wbaker85/tacklebox/pkg/models"
)

// HookRecordModel comment
type HookRecordModel struct {
	DB *sql.DB
}

// Insert is for adding one hook record to the database
func (m *HookRecordModel) Insert(binID, hookID string) (int, error) {
	stmt := `INSERT INTO records (bin_id, hook_id, created)
					 VALUES ($1, $2, CURRENT_TIMESTAMP)
					 RETURNING id`

	var createdID int
	err := m.DB.QueryRow(stmt, binID, hookID).Scan(&createdID)
	if err != nil {
		var pgError *pq.Error
		if errors.As(err, &pgError) {
			if pgError.Code.Name() == "foreign_key_violation" {
				return 0, models.ErrInvalidBin
			}
		}
		return 0, err
	}

	return createdID, nil
}

// Destroy is for deleting a record from the database
func (m *HookRecordModel) Destroy(hookID string) error {
	stmt := `DELETE FROM records WHERE hook_id = $1`

	_, err := m.DB.Exec(stmt, hookID)
	if err != nil {
		return err
	}

	return nil
}

// Get a slice containing all of the hooks for the provided binID belonging to the userID
func (m *HookRecordModel) Get(binID string) ([]*models.HookRecord, error) {
	stmt := `SELECT id, bin_id, hook_id, created
					 FROM records
					 WHERE bin_id = $1`

	rows, err := m.DB.Query(stmt, binID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hooks := []*models.HookRecord{}
	for rows.Next() {
		h := &models.HookRecord{}
		err = rows.Scan(&h.ID, &h.BinID, &h.HookID, &h.Created)
		if err != nil {
			return nil, err
		}

		hooks = append(hooks, h)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return hooks, nil
}

// GetOne returns a single hook record, given a hook ID
func (m *HookRecordModel) GetOne(hookID string) (*models.HookRecord, error) {
	stmt := `SELECT bin_id, hook_id, created
					 FROM records
					 WHERE hook_id = $1`

	row := m.DB.QueryRow(stmt, hookID)

	res := &models.HookRecord{}
	err := row.Scan(&res.BinID, &res.HookID, &res.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrInvalidHook
		}
		return nil, err
	}

	return res, nil
}

// CheckBinOwnership verifies that a provider userID has access to a binID for deleting, reading, etc.
func (m *HookRecordModel) CheckBinOwnership(userID int, binID string) (bool, error) {
	stmt := `SELECT user_id
					 FROM bins
					 WHERE id = $1`

	row := m.DB.QueryRow(stmt, binID)

	var res int
	err := row.Scan(&res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, models.ErrInvalidBin
		}
		return false, err
	}

	return res == userID, nil
}

// CheckRecordOwnership verifies that a hook record belongs to the given user
func (m *HookRecordModel) CheckRecordOwnership(userID int, hookID string) (bool, error) {
	stmt := `SELECT b.user_id
					 FROM records AS r
					 INNER JOIN bins AS b ON r.bin_id = b.id
					 WHERE r.hook_id = $1`

	row := m.DB.QueryRow(stmt, hookID)

	var res int
	err := row.Scan(&res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, models.ErrInvalidHook
		}
		return false, err
	}

	return res == userID, nil
}
