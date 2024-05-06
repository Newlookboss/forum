package mocks

import (
	"github.com/gofrs/uuid"
	"net/http"
)

type SessionModel struct{}

func (m *SessionModel) AddToken(id int, token string) error {
	// Реализация моковой функции добавления токена
	return nil
}

func (m *SessionModel) GetToken(token string) (string, error) {
	// Реализация моковой функции получения токена
	return "", nil
}

func (m *SessionModel) GetNameByToken(token string) (string, error) {
	// Реализация моковой функции получения имени по токену
	return "", nil
}

func (m *SessionModel) RemoveToken(token string) error {
	// Реализация моковой функции удаления токена
	return nil
}

func (m *SessionModel) IsExpired(token string) bool {
	// Реализация моковой функции проверки на истечение срока действия токена
	return false
}

func IsLogged(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	var data bool

	if err != nil || cookie.Value == "" {
		data = false
	} else {
		data = true
	}
	return data
}

func GetUUIDToken() string {
	token, err := uuid.NewV4()
	if err != nil {
		// Обработка ошибки, если не удалось сгенерировать UUID
		return ""
	}

	return token.String()
}

func DeleteCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
		// Raw:      "",
	}
	http.SetCookie(w, cookie)
}
