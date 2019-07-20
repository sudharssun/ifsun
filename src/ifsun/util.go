package main

import (
	"fmt"
	"os"
	"time"
)

func exitIfError(err error) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func CurrentDateWithOffset(offsetFromToday int) string {
    currentTime := time.Now()
    offsetTime := currentTime.AddDate(0, 0, offsetFromToday)
    return offsetTime.Format("2006-01-02")        
}