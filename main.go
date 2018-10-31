package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	startTime := 1362984425

	_, err := time.Parse(timeFormat, start)
	if err != nil {
		return errors.New("Start Time is invalid")
	}
	t, _ := time.Parse(timeFormat.String(), "1362984425")
	fmt.Println(timeFormat.Format(time.RFC1123))
	fmt.Println(t)
}
