package fim

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"io"
	"io/fs"
	"log"
	"m-faheem-khan/file-integrity-monitoring/pkg/db"
	"m-faheem-khan/file-integrity-monitoring/pkg/enums"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var symlinkTargets = make(map[string]string)
var mu sync.Mutex

func BuildHashDB(rootDir string, sdb *sql.DB) {
	fmt.Printf("Building Hash DB for %s\n", rootDir)

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Generate hash for files
		if !d.IsDir() && d.Type() != fs.ModeSymlink {
			hash, err := generateSHA256Hash(path)
			if err != nil {
				log.Printf("Error generating hash for %s: %v\n", path, err)
				return nil
			}
			fmt.Printf("%s:\t%x\n", path, hash)
			r := db.Row{
				FilePath:             path,
				ShaHash:              string(hash),
				IntegrityStatus:      enums.INTEGRITY_INITIAL_SCAN,
				LastIntegritScanTime: time.Now(),
				LastEventName:        enums.EVENT_INITIAL_SCAN,
				LastEventNameTime:    time.Now(),
			}

			db.Insert(sdb, r) // insert into db
		}

		// Iterate over Symlink directories
		if d.Type() == fs.ModeSymlink {
			fmt.Printf("%s is a symlink.\n", path)
			// Resolve the symbolic link to get the actual file path
			targetPath, err := os.Readlink(path)
			if err != nil {
				log.Printf("Error reading symbolic link %s: %v\n", path, err)
				return nil
			}

			// Handle cases where targetPath does not start with "/"
			if !strings.HasPrefix(targetPath, "/") {
				targetPath = filepath.Join(filepath.Dir(path), targetPath)
			}

			// Check if the targetPath has already been processed
			mu.Lock()
			_, exists := symlinkTargets[targetPath]
			mu.Unlock()

			if !exists {
				mu.Lock()
				symlinkTargets[targetPath] = targetPath
				mu.Unlock()

				BuildHashDB(targetPath, sdb) // recursive call

			} else {
				log.Printf("Skipping duplicate symlink target: %s\n", targetPath)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error building hash db for %s: %v\n", rootDir, err)
		os.Exit(1)
	}

	fmt.Printf("Finished Generating Hash for %s.\n", rootDir)
}

func generateSHA256Hash(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}
