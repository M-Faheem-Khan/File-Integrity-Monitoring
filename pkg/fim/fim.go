package fim

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"m-faheem-khan/file-integrity-monitoring/pkg/db"
	"m-faheem-khan/file-integrity-monitoring/pkg/enums"
	"os"
	"path/filepath"
	"time"
)

func BuildHashDB(rootDir string, sdb *sql.DB) {
	fmt.Printf("Building Hash DB for %s\n", rootDir)
	count := 0

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Ignore directories, Symlinks & Character Devices
		if !d.IsDir() && d.Type() != fs.ModeSymlink && d.Type() != fs.ModeCharDevice {
			hash, err := generateSHA256Hash(path)
			if err != nil {
				log.Printf("Error generating hash for %s: %v\n", path, err)
				return nil
			}

			r := db.Row{
				FilePath:             path,
				ShaHash:              hash,
				IntegrityStatus:      enums.INTEGRITY_INITIAL_SCAN,
				LastIntegritScanTime: time.Now(),
				LastEventName:        enums.EVENT_INITIAL_SCAN,
				LastEventNameTime:    time.Now(),
			}

			db.Insert(sdb, r) // insert into db
			count++
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error building hash db for %s: %v\n", rootDir, err)
		os.Exit(1)
	}

	fmt.Printf("Finished Generating Hash for %s(%d).\n", rootDir, count)
}

func generateSHA256Hash(filePath string) (string, error) {
	// application is hanging due to large file analysis
	finfo, err := os.Lstat(filePath)
	if err != nil {
		return "", err
	}

	// Lets see what is hanging the program
	fmt.Printf("Generating hash for %s of size %d.\n", filePath, finfo.Size())

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hash_string := hex.EncodeToString(hash.Sum(nil))

	return hash_string, nil
}
