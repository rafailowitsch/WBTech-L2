package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// CurrTime returns current time
func CurrTime() time.Time {
	currTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	return currTime
}
