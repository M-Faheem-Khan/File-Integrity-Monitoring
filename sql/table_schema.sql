CREATE TABLE IF NOT EXISTS FIM_HASHES (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    file_path VARCHAR(4096),
    file_hash VARCHAR(40),
    integrity_status VARCHAR(40),
    last_event VARCHAR(40),
    last_integrity_scan_time DATETIME DEFAULT CURRENT_TIMESTAMP
);
