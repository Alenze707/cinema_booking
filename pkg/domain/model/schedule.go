package model

import "time"

type Schedule struct {
	Moment  time.Time
	MovieID string
	RoomID  string
}
