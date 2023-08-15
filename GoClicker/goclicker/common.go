package goclicker

/*
	Common functions used across multiple
	packages
*/

import (
	"log"
)

// Function that checks for an error
// Returns 0 if there are no errors
// Returns 1 if there are errors
func CheckForError(err error) {
	if err == nil {
		return
	} else {
		log.Fatal("ERROR: ", err)
		return
	}
}
