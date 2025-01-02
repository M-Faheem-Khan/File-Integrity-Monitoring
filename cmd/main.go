package main

import (
	"flag"
	"fmt"
	"m-faheem-khan/file-integrity-monitoring/pkg/db"
	"m-faheem-khan/file-integrity-monitoring/pkg/fim"
)

// Requirements
// - Given a directory
//   - Iterate over all files
//   - Iterate over all subdirectories
//   - Iterate over all symlinks

func printHelp() {
	println("File Integrity Monitoring(FIM)")
	fmt.Printf("--help\tPrints this help page.\n")
	fmt.Printf("--build-hash-db\tBuild database with hashes for all files.\n")
	fmt.Printf("--dir\tDirectory for which hashes must be genrated. Default: /usr/bin \n")
}

func main() {
	// Arguments for the application
	arg_print_help := flag.Bool("help", false, "--help")
	arg_build_hash_db := flag.Bool("build-hash-db", false, "--build-hash-db")
	arg_dir_path := flag.String("dir", "/usr/bin", "--dir=/usr/bin")

	flag.Parse()

	if *arg_print_help {
		printHelp()
	}

	if *arg_build_hash_db {
		sdb := db.GetDatabase()
		fim.BuildHashDB(*arg_dir_path, sdb)
		db.GetDatabase().Close() // close db building hash
	}

}
