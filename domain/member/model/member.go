package model

import "time"

type Member struct {
	Code               string
	Name               string
	DatePenaltyExpired time.Time
}
