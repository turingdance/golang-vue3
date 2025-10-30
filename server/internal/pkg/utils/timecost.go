package utils

import (
	"fmt"
	"time"
)

// @brief：耗时统计函数
// defer timeCost()()
func TimeCost(tags string, params ...interface{}) func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		tc := time.Since(start)
		fmt.Printf(tags+"  "+fmt.Sprintf(" timecost = %dms", tc/(1000*1000))+"  ", params...)
		return tc
	}
}
