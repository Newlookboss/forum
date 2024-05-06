package mocks

import "mkassymk/forum/internal/models"

type LikeModel struct{}

func (m *LikeModel) LikeInsert(likeData models.UserLikeData, likedByUserID int) error {
	// Здесь можно просто вернуть nil или фиктивное значение ошибки
	return nil
}

func (m *LikeModel) DislikeInsert(dislikeData models.UserDislikeData, dislikedByUserID int) error {
	// Здесь можно просто вернуть nil или фиктивное значение ошибки
	return nil
}

func (m *LikeModel) RemoveLike(postid int, likedByUserID int) error {
	// Здесь можно просто вернуть nil или фиктивное значение ошибки
	return nil
}

func (m *LikeModel) RemoveDislike(postid int, likedByUserID int) error {
	// Здесь можно просто вернуть nil или фиктивное значение ошибки
	return nil
}

func (m *LikeModel) IsLikedByUser(user int, postid int) bool {
	// Здесь можно просто вернуть true или false
	return true
}

func (m *LikeModel) IsDislikedByUser(user int, postid int) bool {
	// Здесь можно просто вернуть true или false
	return true
}

func (m *LikeModel) UpdateLikesCount() error {
	// Здесь можно просто вернуть nil или фиктивное значение ошибки
	return nil
}
