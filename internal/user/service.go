package user

import (
	"errors"

	"github.com/MortalSC/GoForIM/pkg/db"
)

func RegisterUser(user *User) error {
	if err := user.HashPassword(); err != nil {
		return err
	}
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := db.DB.Exec(query, user.Username, user.Password)
	return err
}

func AuthenticateUser(username, password string) (*User, error) {
	var user User
	query := "SELECT id, username, password FROM users WHERE username = ?"
	err := db.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if !user.CheckPassword(password) {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}
