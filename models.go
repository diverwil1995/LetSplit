package main

import (
	"crypto/sha256"
	"encoding"
	"errors"
	"time"
	"unicode"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	// Id       int    `json:"id"`
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
func (u *User) GenerateUuid() {
	u.Uuid = uuid.NewV4().String()
}

type Group struct {
	// Id     int    `json:"id"`
	Uuid     string `json:"uuid"`
	UserUuid string `json:"user_uuid"`
	Title    string `json:"title"`
}

var (
	ErrUserUuidEmpty = errors.New("user_uuid can't be empty")
	ErrTitleEmpty    = errors.New("title can't be empty")
)

func (g *Group) Validate() error {
	if g == nil {
		return nil
	}
	if g.UserUuid == "" {
		return ErrUserUuidEmpty
	}
	if g.Title == "" {
		return ErrTitleEmpty
	}
	return nil
}
func (g *Group) GenerateUuid() {
	g.Uuid = uuid.NewV4().String()
}

// TODO: Debt field defult value 0
type Member struct {
	// Id      int     `json:"id"`
	Uuid      string  `json:"uuid"`
	GroupUuid string  `json:"group_uuid"`
	Name      string  `json:"name"`
	Debt      float64 `json:"debt"`
}

var (
	ErrGroupUuidEmpty  = errors.New("group_uuid can't be empty")
	ErrMemberNameEmpty = errors.New("member name can't be empty")
)

func (m *Member) Validate() error {
	if m == nil {
		return nil
	}
	if m.GroupUuid == "" {
		return ErrGroupUuidEmpty
	}
	if m.Name == "" {
		return ErrMemberNameEmpty
	}
	m.Debt = 0
	return nil
}
func (m *Member) GenerateUuid() {
	m.Uuid = uuid.NewV4().String()
}

type Expense struct {
	// Id             int       `json:"id"`
	Uuid             string    `json:"uuid"`
	GroupUuid        string    `json:"group_uuid"`
	PayerUuid        string    `json:"payer_uuid"`
	ParticipantsUuid []string  `json:"participants_uuid"`
	Item             string    `json:"item"`
	Date             time.Time `json:"date"`
	Cost             float64   `json:"cost"`
}

var (
	ErrPayerUuidEmpty    = errors.New("payer_uuid can't be empty")
	ErrParticipantsEmpty = errors.New("participants_uuid can't be empty")
	ErrItemEmpty         = errors.New("item can't be empty")
	ErrDateEmpty         = errors.New("date can't be empty")
	ErrCostEmpty         = errors.New("cost can't be empty")
)

func (e *Expense) Validate() error {
	if e == nil {
		return nil
	}
	if e.GroupUuid == "" {
		return ErrGroupUuidEmpty
	}
	if e.PayerUuid == "" {
		return ErrPayerUuidEmpty
	}
	if len(e.ParticipantsUuid) == 0 {
		return ErrParticipantsEmpty
	}
	if e.Item == "" {
		return ErrItemEmpty
	}
	if e.Date.Format("2006-01-02") == "" {
		return ErrDateEmpty
	}
	if e.Cost <= 0 {
		return ErrCostEmpty
	}
	return nil
}
func (e *Expense) GenerateUuid() {
	e.Uuid = uuid.NewV4().String()
}
