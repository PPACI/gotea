package libs

import (
	"fmt"
	"time"
)

func StartTimer(duration uint16) {
	if duration <= 0 {
		Done()
		return
	}
	ticker := time.NewTicker(1 * time.Minute)
	var timeTracker uint16 = 0
	for timeTracker < duration {
		<-ticker.C
		timeTracker++
		fmt.Printf("Time elapsed : %d minutes\n", timeTracker)
	}
	Done()
}
