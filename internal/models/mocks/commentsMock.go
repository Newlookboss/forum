package mocks

import "mkassymk/forum/internal/models"

type CommentModel struct{}

func (m *CommentModel) CommentInsert(comment string, authorID int, postId int) error {
	// Реализация моковой функции вставки комментария
	return nil
}

func (m *CommentModel) UpdateCommentsCount() error {
	// Реализация моковой функции обновления счетчика комментариев
	return nil
}

func (m *CommentModel) GetComments(user models.UserModelInterface, postId int, authorID int) ([]*models.PostComments, error) {
	// Реализация моковой функции получения комментариев
	return nil, nil
}

func (m *CommentModel) CommentLikeInsert(likeData models.CommentLikeData, likedBy int) error {
	// Реализация моковой функции лайка комментария
	return nil
}

func (m *CommentModel) RemoveCommentLike(commentid int, likedBy int) error {
	// Реализация моковой функции удаления лайка комментария
	return nil
}

func (m *CommentModel) IsCommentLikedByUser(user int, commentid int) bool {
	// Реализация моковой функции проверки, лайкнут ли комментарий пользователем
	return false
}

func (m *CommentModel) CommentDislikeInsert(dislikeData models.CommentDislikeData, dislikedBy int) error {
	// Реализация моковой функции дизлайка комментария
	return nil
}

func (m *CommentModel) RemoveCommentDislike(commentid int, likedBy int) error {
	// Реализация моковой функции удаления дизлайка комментария
	return nil
}

func (m *CommentModel) IsCommentDislikedByUser(userID int, commentid int) bool {
	// Реализация моковой функции проверки, дизлайкнут ли комментарий пользователем
	return false
}

func (m *CommentModel) UpdateCommentLikesCount() error {
	// Реализация моковой функции обновления счетчика лайков и дизлайков комментариев
	return nil
}
