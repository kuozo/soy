package pkg

import (
	"time"
)

// GetCurrentTime get current time
func GetCurrentTime() string {

	return time.Now().Format(time.RFC822)
}
