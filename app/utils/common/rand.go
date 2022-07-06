package common

import (
	"math"
	"math/rand"
	"time"
)

func RandNum(digits int) int {
	rand.Seed(time.Now().UnixNano())
	max := math.Pow(float64(10), float64(digits))
	return rand.Intn(int(max))
}
