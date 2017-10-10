package meuponto

import "time"

type Entry struct {
	Comment string    `json:"comment"`
	Time    time.Time `json:"time"`
}
