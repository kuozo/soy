package handler

import (
	"testing"
)

func TestCalDays(t *testing.T) {
	minutes := calMinutes()
	days := calDays(minutes)
	t.Logf("days: %d", days)
}
