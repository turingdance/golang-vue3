package utils

import (
	"fmt"
	"strconv"
)

func Atoi[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint | int | uint32 | uint64](input string) T {
	t, _ := (strconv.Atoi(input))
	return T(t)
}
func ItoA(input any) string {
	return fmt.Sprintf("%d", input)
}
