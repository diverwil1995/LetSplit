package main

import "time"

type Group struct {
	Id      int
	Uuid    string
	Name    string
	Members []*Member
	Records []*Record
	Report  *Report
}
type Member struct {
	Id   int
	Uuid string
	Name string
}
type Expense struct {
	Id     int
	Uuid   string
	Date   time.Time
	What   string
	Amount float64
	By     *Member
	For    []*Member
	Note   string
}
type Record struct {
	Id       int
	Uuid     string
	Expenses []*Expense
	DebtList []*Debt
}
type Debt struct {
}
type Report struct {
}
