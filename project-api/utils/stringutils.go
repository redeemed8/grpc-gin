package utils

import (
	"golang.org/x/exp/rand"
	"strconv"
	"time"
)

func MakeCodeWithNumber(length int, symbol int) (ret string) {
	if length <= 0 {
		return ""
	}
	if length > 25 {
		length = 25
	}
	rng := rand.New(rand.NewSource(uint64(time.Now().UnixNano() + int64(symbol))))
	for i := 0; i < length; i++ {
		num := rng.Intn(10)
		ret += strconv.Itoa(num)
	}
	return ret
}
