package mocks

import (
	"mkassymk/forum/internal/models"
	"time"
)

type PostModel struct{}

var mockPost = &models.Post{
	PostID:   1,
	Author:   "Matsuo Bashō",
	Title:    "An old silent pond",
	Content:  "An old silent pond...",
	Created:  time.Now(),
	Likes:    0,
	Dislikes: 0,
	Comments: 0,
	Tags:     "Other",
}

func (m *PostModel) Insert(userID int, author string, title string, content string, tags []string) (int, error) {
	// Здесь можно вернуть фиктивный ID поста или ошибку
	return 1, nil
}

func (m *PostModel) CategoryInsert(postid int64, categories []string) error {
	// Здесь можно просто вернуть nil или фиктивное значение ошибки
	return nil
}

func (m *PostModel) GetPostbyPostID(id int) (*models.Post, error) {
	// Здесь можно вернуть фиктивный пост или ошибку
	if id != 1 {
		return nil, models.ErrNoRecord
	}
	return mockPost, nil
}

func (m *PostModel) GetUserPosts(id int) (map[int]*models.Post, error) {
	// Здесь можно вернуть фиктивные посты пользователя или ошибку
	posts := map[int]*models.Post{1: mockPost} // Пример
	return posts, nil
}

func (m *PostModel) GetUserLikedPosts(userID int) (map[int]*models.Post, error) {
	// Здесь можно вернуть фиктивные посты, которые пользователь лайкнул, или ошибку
	likedPosts := map[int]*models.Post{1: mockPost} // Пример
	return likedPosts, nil
}

func (m *PostModel) Latest(RM models.LikeModelInterface) (map[int]*models.Post, error) {
	// Здесь можно вернуть фиктивные последние посты или ошибку
	latestPosts := map[int]*models.Post{1: mockPost} // Пример
	return latestPosts, nil
}

func (m *PostModel) FilteredPosts(categories []string) (map[int]*models.Post, error) {
	// Здесь можно вернуть фиктивные посты, отфильтрованные по категориям, или ошибку
	filteredPosts := map[int]*models.Post{1: mockPost} // Пример
	return filteredPosts, nil
}
