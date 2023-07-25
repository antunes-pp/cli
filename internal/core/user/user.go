package user

import (
	"errors"
	"net/mail"
)

type user struct {
	Name   string   `json:"user_name"`
	Email  string   `json:"user_email"`
	Squads []string `json:"squads"`
}

func NewUser(name, email string, squads []string) *user {
	u := user{
		Name:   name,
		Email:  email,
		Squads: squads,
	}

	return &u
}

func (u *user) Validate() error {
	if u.Name == "" {
		return errors.New("name_is_required")
	}

	if u.Email == "" {
		return errors.New("email_is_required")
	}

	if len(u.Squads) == 0 {
		return errors.New("squads_is_empty")
	}

	adds, err := mail.ParseAddress(u.Email)

	if err != nil {
		return errors.New("invalid_address")
	}

	u.Email = adds.Address

	return nil
}
