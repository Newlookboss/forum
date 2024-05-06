package models

import (
	"database/sql"
	"errors"
	"html"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    int
	Firstname string
	Lastname  string
	Email     string
	Token     string
	Expiry    time.Time
}

type UserModel struct {
	DB *sql.DB
}

type UserModelInterface interface {
	InsertUser(firstname, lastname, email, password string) error
	GetUserNamebyUserID(id int) (string, error)
	GetUserInfo(r *http.Request) (*User, error)
	Authenticate(email, password string) (int, error)
}

func (u *UserModel) InsertUser(firstname, lastname, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	firstname = html.EscapeString(firstname)
	lastname = html.EscapeString(lastname)

	stmt := `INSERT INTO users (firstname, lastname, email, hashedPassword) VALUES (?, ?, ?, ?)`

	_, err = u.DB.Exec(stmt, firstname, lastname, email, string(hashedPassword))

	if err != nil {
		return err
	}

	return nil
}

func (u *UserModel) GetUserNamebyUserID(id int) (string, error) {
	stmt := `SELECT firstname, lastname FROM users
    WHERE userID = ?`

	row := u.DB.QueryRow(stmt, id)
	firstname := ""
	lastname := ""
	err := row.Scan(&firstname, &lastname)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "undefined", ErrNoRecord
		} else {
			return "undefined", err
		}
	}
	return firstname + " " + lastname, nil
}

func (u *UserModel) GetUserInfo(r *http.Request) (*User, error) {
	user := &User{}

	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, err
	}
	token := cookie.Value

	stmt := `SELECT userID, firstname, lastname, email, token, expiry FROM users
    WHERE token = ?`

	row := u.DB.QueryRow(stmt, token)

	err = row.Scan(&user.UserID, &user.Firstname, &user.Lastname, &user.Email, &user.Token, &user.Expiry)

	return user, nil
}

func (u *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := `SELECT userID, hashedPassword FROM users WHERE email = ?`

	err := u.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}
