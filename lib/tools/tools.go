package tools

import (
	"fmt"
	"time"
)

func GetCronTime(timeString string, week int) string {
	t, _ := time.Parse("15:04:05", timeString)
	h := fmt.Sprintf("%d %d %d * * %d", t.Second(), t.Minute(), t.Hour(), week)
	return h
}
