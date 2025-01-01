# File-Integrity-Monitoring (UNDER-DEVELOPMENT)
A File Integrity Monitoring solution written in Golang.

[![Scorecard supply-chain security](https://github.com/M-Faheem-Khan/File-Integrity-Monitoring/actions/workflows/scorecard.yml/badge.svg)](https://github.com/M-Faheem-Khan/File-Integrity-Monitoring/actions/workflows/scorecard.yml)
[![golangci-lint](https://github.com/M-Faheem-Khan/File-Integrity-Monitoring/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/M-Faheem-Khan/File-Integrity-Monitoring/actions/workflows/golangci-lint.yml)
[![CodeQL Advanced](https://github.com/M-Faheem-Khan/File-Integrity-Monitoring/actions/workflows/codeql.yml/badge.svg)](https://github.com/M-Faheem-Khan/File-Integrity-Monitoring/actions/workflows/codeql.yml)
[![Dependency Review](https://github.com/M-Faheem-Khan/File-Integrity-Monitoring/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/M-Faheem-Khan/File-Integrity-Monitoring/actions/workflows/dependency-review.yml)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/9877/badge)](https://www.bestpractices.dev/projects/9877)
----

The purpose of the tool is to identify and detect changes made to any files being 'watched'. The following events
are looked for by the tool <code>create</code>, <code>modify</code> & <code>delete</code>. The before the tool can
begin 'watching' a directory it needs to build a hash database(<code>Sqllite3</code>) of all the existing files. The application runs in two modes <code>kernel</code> & <code>cron</code>. Kernel mode monitors the system calls made
related to the files in the 'watch' directory. Cron mode runs generates hashes for all the files in the 'watch'
directory. Both modes use hashes to keep track of integrity.

**Why use Cron mode?**
Cron mode should be used if you only want to monitor the files that are changed, real-time notification is not a requirement.
The changes include <code>create</code>, <code>modify</code> & <code>delete</code>.

**Why use Kernel mode?**
Kernel mode has the added benfit that application will be able to report the integrity violaiton in near real-time. Allowing
you to view when, who & what changes were made.



- Program must take the following arguments:
    - <code>--build-hash-db</code>: Generate hash for given directory.
    - <code>--dir</code>: Directory for which hash must be generated/monitored.

- Generate hash for all files in a given directory
    - Generate hash for all symlinks
    - Symlinks must only be scanned once.
