# Database Design

**Table Of Contents**
* [Tamper Protection](#tamper-protection)
  + [Encryption](#encryption)
  + [Permissions](#permissions)

As the application needs a local database to keep track of values. SQLlite will be used due to portability, simplicity and reliability. A single table will be created <code>FIM_HASH</code> with the below schema and only <code>INSERT</code> & <code>UPDATE</code> operations are permitted by the application.

```SQL
CREATE TABLE IF NOT EXISTS FIM_HASHES (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    file_path VARCHAR(4096),      -- max linux file path size
    file_hash VARCHAR(40),        -- SHA256 hash size(hex)
    integrity_status VARCHAR(40), -- Integrity Status
    last_event VARCHAR(40),       -- Event Name
    last_integrity_scan_time DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## Tamper Protection
In order to prevent the database from being tampered with the following mechanisms are implemented:
- Database is encrypted at rest.
- Database encryption key is only provided to the application.
- At runtime the database can only be read/write/executed by the user executing the application.

### Encryption
### Permissions
