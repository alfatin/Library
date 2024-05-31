package model

import "time"

type List struct {
	Quantity int
	Book []Book
}

type Book struct {
	Code         string
	Title        string
	Author       string
	Stock        int
	BorrowedBy   string
	DateBorrowed *time.Time
}
