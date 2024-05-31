package model

import "time"

type Member struct {
	Code               string
	Name               string
	DatePenaltyExpired time.Time
	Quantity           int
}

type List struct {
	Name     string
	Quantity int
}

