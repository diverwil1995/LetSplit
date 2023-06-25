package main

import "time"

type User struct {
	Id       int    `json:"id"`
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"-"`
}
type Group struct {
	Id     int    `json:"id"`
	Uuid   string `json:"uuid"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
}
type Member struct {
	Id      int     `json:"id"`
	Uuid    string  `json:"uuid"`
	GroupId int     `json:"group_id"`
	Name    string  `json:"name"`
	Debt    float64 `json:"debt"`
}
type Expense struct {
	Id             int       `json:"id"`
	Uuid           string    `json:"uuid"`
	What           string    `json:"what"`
	Date           time.Time `json:"date"`
	Cost           float64   `json:"cost"`
	PayerId        int       `json:"payer_id"`
	ParticipantsId []int     `json:"participants_id"`
}
