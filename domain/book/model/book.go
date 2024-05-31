package model

import "time"

type Book struct {
	Code         string
	Title        string
	Author       string
	Stock        int
	BorrowedBy   string
	DateBorrowed time.Time
}
