package postgres

import (
	"database/sql"
)

// HookRecordModel comment
type HookRecordModel struct {
	DB *sql.DB
}

// InsertOne is for adding one hook record to the database
func (m *HookRecordModel) InsertOne(binID, hookID string) error {
	stmt := `INSERT INTO records (bin_id, hook_id, date) VALUES ($1, $2, CURRENT_TIMESTAMP);`

	_, err := m.DB.Exec(stmt, binID, hookID)
	if err != nil {
		return err
	}

	return nil
}
