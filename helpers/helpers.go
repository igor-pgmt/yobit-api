package helpers

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// FormatFloat returns human-readable price
func FormatFloat(f float64) string {
	return fmt.Sprint(strconv.FormatFloat(f, 'f', -1, 64))
}

// Percent returns percentage of the number
func Percent(pct float64, number float64) float64 {
	percent := ((number * pct) / float64(100))
	return percent
}

// PercentOf returns percent from number
func PercentOf(current float64, all float64) float64 {
	percent := (current * float64(100)) / all
	return percent
}

// PercentChange returns growth\fall percentage value
func PercentChange(before float64, after float64) float64 {
	diff := after - before
	realDiff := diff / before
	percentDiff := 100 * realDiff

	return percentDiff
}

// RandomDelay returns a random number on the specified interval
func RandomDelay(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max+1-min) + min
}
