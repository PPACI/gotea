package libs

import (
	"time"
	"fmt"
	"github.com/sqweek/dialog"
)

func StartTimer(duration uint16) {
	if duration <=0 {
		fmt.Println("Timer must be > 0")
		return
	}
	ticker := time.NewTicker(1 * time.Minute)
	var timeTracker uint16 = 0
	for timeTracker < duration {
		_ <- ticker.C
		timeTracker ++
		fmt.Printf("Time elapsed : %d minutes", timeTracker)
	}
	dialog.Message("Done !").Title("Tea Timer").Info()
}