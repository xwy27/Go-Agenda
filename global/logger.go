package global

import (
	"fmt"
	"log"
	"os"
)

// ErrorLog is the logger to output the error
// info
// example: ErrorLog.Println(...)
var errorLog = log.New(os.Stderr, "", 0)

// PrintError is useful
func PrintError(err error, success string) {
	if err != nil {
		errorLog.Println(err.Error())
		return
	}
	if len(success) != 0 {
		fmt.Println(success)
	}
}
