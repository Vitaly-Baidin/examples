package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddUpdatedAt, downAddUpdatedAt)
}

func upAddUpdatedAt(tx *sql.Tx) error {
	_, err := tx.Exec(`ALTER TABLE users ADD COLUMN IF NOT EXISTS update_at TIMESTAMP DEFAULT now();`)

	if err != nil {
		return err
	}

	return nil
}

func downAddUpdatedAt(tx *sql.Tx) error {
	_, err := tx.Exec(`ALTER TABLE users DROP COLUMN IF EXISTS update_at;`)
	if err != nil {
		return err
	}
	return nil
}
