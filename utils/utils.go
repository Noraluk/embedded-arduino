package utils

import "time"

func Delay(t uint32) {
	time.Sleep(time.Duration(1000000 * t))
}
