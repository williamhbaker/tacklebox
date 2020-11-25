package postgres

import (
	"database/sql"
)

// HookRecordModel comment
type HookRecordModel struct {
	DB *sql.DB
}

// InsertOne is for adding one hook record to the database
func (m *HookRecordModel) InsertOne(binID, hookID string) (int, error) {
	stmt := `INSERT INTO records VALUES(bin_id, hook_id) ($1, $2)`

	result, err := m.DB.Exec(stmt, binID, hookID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
