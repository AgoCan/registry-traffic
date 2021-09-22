package main

import "fmt"

var (
	Version  string
	CommitID string
)

func showVersion() {
	fmt.Printf("Version: %s, commit_id: %s", Version, CommitID)
}
