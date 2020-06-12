package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	timeZone = "Asia/Shanghai"
)

// Girl for girl entity
type Girl struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Days int    `json:"days"`
}

func location() (*time.Location, error) {
	return time.LoadLocation(timeZone)
}

func birthDay() (time.Time, error) {
	loc, err := location()
	if err != nil {
		return time.Now(), err
	}
	brith := time.Date(2020, time.May, 27, 16, 40, 0, 0, loc)
	return brith, nil
}

func today() (time.Time, error) {
	today := time.Now()
	loc, err := location()
	if err != nil {
		return today, err
	}
	return today.In(loc), nil
}

func calMinutes() float64 {
	td, err := today()
	if err != nil {
		return 0
	}
	bd, err := birthDay()
	if err != nil {
		return 0
	}
	return td.Sub(bd).Minutes()
}

func calDays(minutes float64) int {
	return int(minutes / 60 / 24)
}

func calYear(minutes float64) int {
	return int(minutes / 60 / 24 / 365)
}

// RHandler for record my girl days
func RHandler(w http.ResponseWriter, r *http.Request) {
	minutes := calMinutes()
	if minutes == 0 {
		w.Write([]byte("some error happend"))
		return
	}
	girl := Girl{
		Name: "zxr",
		Age:  calYear(minutes),
		Days: calDays(minutes),
	}
	json.NewEncoder(w).Encode(girl)
}
