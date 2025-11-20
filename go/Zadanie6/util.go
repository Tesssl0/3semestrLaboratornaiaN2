// util.go
package main

import (
	"math/rand"
	"strconv"
	"time"
)

type RandomGenerator struct{}

func (rg *RandomGenerator) GenerateRandomNumbers(count int, minVal int, maxVal int) *Array {
	numbers := NewArray(count)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		num := rand.Intn(maxVal-minVal+1) + minVal
		numbers.Add(strconv.Itoa(num))
	}

	return numbers
}

type Timer struct {
	startTime time.Time
}

func (t *Timer) Start() {
	t.startTime = time.Now()
}

func (t *Timer) ElapsedMillis() int64 {
	return time.Since(t.startTime).Milliseconds()
}
