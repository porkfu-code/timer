package timer

import (
	"context"
	"fmt"
	"time"
)

/*
Just show a simple timer on the screen that counts down until 0.

The display of the timer shall be like:
HH:MM:SS

Optional
- The timer binary should allow user args like the start time, run like:
  ~ % ./timer -H 1 -M 30 -S 59  # 1 hour, 30 minute, 59 second timer

- If the user does something like specify 61 minutes, this should be displayed as a 1 hour, 1 minute
  timer; i.e. starts from 01:01:00
*/

type Timer struct {
	tEnd  time.Time
	tTick time.Duration
}

func NewTimer(duration time.Duration, tickTime time.Duration) *Timer {
	return &Timer{
		tEnd:  time.Now().Add(duration),
		tTick: tickTime,
	}
}

func (t *Timer) displayRemainingTime() {
	remaining := time.Until(t.tEnd)
	if remaining < 0 {
		remaining = 0
	}

	totalSeconds := int(remaining.Seconds())
	h := totalSeconds / 3600
	m := (totalSeconds % 3600) / 60
	s := totalSeconds % 60

	fmt.Printf("\rRemaining time: %02d:%02d:%02d", h, m, s)
}

func (t *Timer) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nTimer stopped.")
			return
		default:
			t.displayRemainingTime()
			if time.Now().After(t.tEnd) {
				fmt.Println("\nDone!")
				return
			}
			time.Sleep(t.tTick)
		}
	}
}
