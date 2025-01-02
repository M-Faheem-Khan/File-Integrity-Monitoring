package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"m-faheem-khan/file-integrity-monitoring/pkg/enums"
	"time"

	_ "modernc.org/sqlite"
)

type Row struct {
	FilePath             string
	ShaHash              string
	IntegrityStatus      enums.Integrity
	LastIntegritScanTime time.Time
	LastEventName        enums.Event
	LastEventNameTime    time.Time
}

func GetDatabase() *sql.DB {
	dsn := fmt.Sprintf("file:%s.db?mode=rwc&kdf_iter=100000&kdf_keylen=32&cipher=aes-256-gcm&cipher_params=iv=auto", "fim")

	// Open the SQLite database with encryption
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS FIM_HASHES (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		file_path VARCHAR(4096),
		file_hash VARCHAR(40),
		integrity_status VARCHAR(40),
		last_integrity_scan_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		last_event VARCHAR(40),
		last_event_scan_time DATETIME DEFAULT CURRENT_TIMESTAMP);`)

	if err != nil {
		panic(err)
	}
	return db
}

func Insert(db *sql.DB, row Row) bool {
	_, err := db.Exec(
		"INSERT INTO FIM_HASHES (file_path, file_hash, integrity_status, last_event) VALUES (?, ?, ?, ?);",
		row.FilePath, row.ShaHash, row.IntegrityStatus.String(), row.LastEventName.String())

	if err != nil {
		slog.Error(err.Error())
		return false
	}
	return true
}

func Update(db *sql.DB, row Row) bool {
	// Insert a Row
	_, err := db.Exec(
		"UPDATE FIM_HASHES SET file_hash=?, integrity_status=?, last_integrity_scan_time=?, last_event=?, last_event_scan_time=?",
		row.ShaHash,
	)

	if err != nil {
		slog.Error(err.Error())
		return false
	}
	return true
}
