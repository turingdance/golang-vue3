package utils

import (
	"fmt"
	"time"
)

var currentIndex = 0

func BillNo() string {
	timeNow := time.Now()
	currentIndex = currentIndex + 1
	str := timeNow.Format("20060102150405")
	return fmt.Sprintf("%s%02d", str, currentIndex)
}
