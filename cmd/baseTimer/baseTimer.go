package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/porkfu-code/timer"
)

func main() {
	h := flag.Int("H", 0, "Hours")
	m := flag.Int("M", 0, "Minutes")
	s := flag.Int("S", 0, "Seconds")
	flag.Parse()

	// Logic: 61 minutes automatically becomes 1h 1m
	totalDuration := time.Duration(*h)*time.Hour +
		time.Duration(*m)*time.Minute +
		time.Duration(*s)*time.Second

	if totalDuration <= 0 {
		fmt.Println("Please specify a duration greater than 0.")
		return
	}

	timer := timer.NewTimer(totalDuration, 1*time.Second)
	timer.Run(context.Background())
}
