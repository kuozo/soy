package handler

import (
	"math"
	"testing"
)

func TestCalDays(t *testing.T) {
	minutes := calMinutes()
	days := calDays(minutes)
	t.Logf("days: %d", days)
}

func TestCalSome(t *testing.T) {
	minutes := calMinutes()
	t.Logf("%f", math.Ceil(minutes/60/24))
}
