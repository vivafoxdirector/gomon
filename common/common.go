package common

import "time"

type Record struct {
	Code  string
	Time  time.Time
	Value map[string]string
}
