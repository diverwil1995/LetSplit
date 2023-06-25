package main

import (
	"crypto/sha256"
	"encoding"
	"errors"
	"time"
	"unicode"
)

type User struct {
	Id       int    `json:"id"`
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"-"`
}

var (
	ErrUsernameEmpty = errors.New("username is empty")
	ErrPasswordEmpty = errors.New("password is empty")
)

func (u *User) Validate() error {
	if u == nil {
		return nil
	}
	if u.Username == "" {
		return ErrUsernameEmpty
	}
	if u.Password == "" {
		return ErrPasswordEmpty
	}
	return nil
}
func (u *User) CheckPassword() (bool, error) {
	const (
		minLenth = 10
		maxLenth = 25
	)
	var (
		ErrTooShort = errors.New("password is too short")
		ErrTooLong  = errors.New("password is too long")
		hasSpecial  = false
		hasUpper    = false
		hasLower    = false
		hasDigit    = false
	)
	if len(u.Password) < minLenth {
		return false, ErrTooShort
	}
	if len(u.Password) > maxLenth {
		return false, ErrTooLong
	}
	for _, v := range u.Password {
		if unicode.IsSymbol(v) {
			hasSpecial = true
		}
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasDigit = true
		}
		if hasSpecial && hasUpper && hasLower && hasDigit {
			break
		}
	}
	return true, nil
}

var (
	ErrPasswordHashFailed = errors.New("password hash failed")
)

func (u User) HashPassword() ([]byte, error) {
	hash := sha256.New()
	hash.Write([]byte(u.Password))
	marshaler, ok := hash.(encoding.BinaryMarshaler)
	if !ok {
		return nil, ErrPasswordHashFailed
	}
	pwd, err := marshaler.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return pwd, nil
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
