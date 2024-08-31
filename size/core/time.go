package core

import (
	"fmt"
	"time"
)

type Timer struct {
	StartTime time.Time
}

func (t *Timer) SetTimer(callback func(int64) float32, timerChan chan any) {
	// timeChan := make(chan any);
	ticker := time.NewTicker(1 * time.Millisecond)

	select {
	case <-ticker.C:
		// fmt.Printf("\r%.2f", callback(t.StartTime.Unix()/int64(time.Millisecond)))
		fmt.Println("Here we got it")
	case <-timerChan:
		ticker.Stop()
	}

}
