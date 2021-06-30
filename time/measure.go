package time

import (
	"fmt"
	"time"
)

func MeasureTime(start time.Time, banner string) {
	end := time.Now()
	fmt.Printf("Target: %s took %.2f seconds\n", banner, end.Sub(start).Seconds())
}
