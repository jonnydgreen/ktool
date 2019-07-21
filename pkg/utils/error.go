package utils

import (
	"log"
)

// ErrorHandler handles errors if they occur
func ErrorHandler(err error) {
	if err != nil {
		log.Fatalf("error occured in ktool: %v", err)
	}
}