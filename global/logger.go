package global

import (
	"log"
	"os"
)

// ErrorLog is the logger to output the error
// info
// example: ErrorLog.Println(...)
var ErrorLog = log.New(os.Stderr, "", 0)
