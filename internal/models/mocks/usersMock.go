package mocks

import (
	"mkassymk/forum/internal/models"
	"net/http"
	"time"
)

type UserModel struct{}

func (u UserModel) InsertUser(firstname, lastname, email, password string) error {
	if email == "dupe@example.com" {
		return models.ErrDuplicateEmail
	}
	return nil
}

func (u UserModel) GetUserNamebyUserID(id int) (string, error) {
	// Здесь можно вернуть фиктивное имя пользователя или ошибку
	return "John Doe", nil
}

func (u UserModel) GetUserInfo(r *http.Request) (*models.User, error) {
	// Здесь можно вернуть фиктивного пользователя или ошибку
	return &models.User{
		UserID:    1,
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		Token:     "abc123",
		Expiry:    time.Now(),
	}, nil
}

func (u UserModel) Authenticate(email, password string) (int, error) {
	// Здесь можно вернуть фиктивный ID пользователя или ошибку
	return 1, nil
}
