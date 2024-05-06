package models

import (
	"database/sql"
	"errors"
)

type CommentLikeData struct {
	PostID    int  `json:"postID"`
	CommentID int  `json:"commentID"`
	Likes     int  `json:"likeCount"`
	IsLiked   bool `json:"isLiked"`
}

type CommentDislikeData struct {
	PostID     int  `json:"postID"`
	CommentID  int  `json:"commentID"`
	Dislikes   int  `json:"dislikeCount"`
	IsDisliked bool `json:"isDisliked"`
}

type PostComments struct {
	PostID     int
	CommentID  int
	Comment    string
	Author     string
	AuthorID   int
	Likes      int
	Dislikes   int
	IsLiked    bool
	IsDisliked bool
}

type CommentModel struct {
	DB *sql.DB
}

type CommentModelInterface interface {
	CommentInsert(comment string, authorID int, postId int) error
	UpdateCommentsCount() error
	GetComments(user UserModelInterface, postId int, authorID int) ([]*PostComments, error)
	CommentLikeInsert(likeData CommentLikeData, likedBy int) error
	RemoveCommentLike(commentid int, likedBy int) error
	IsCommentLikedByUser(user int, commentid int) bool
	CommentDislikeInsert(dislikeData CommentDislikeData, dislikedBy int) error
	RemoveCommentDislike(commentid int, likedBy int) error
	IsCommentDislikedByUser(userID int, commentid int) bool
	UpdateCommentLikesCount() error
}

func (CM *CommentModel) CommentInsert(comment string, authorID int, postId int) error {
	stmt := `INSERT INTO comments (postid, comment, authorID, likes_count, dislikes_count) VALUES(?, ?, ?, '0', '0');`

	_, err := CM.DB.Exec(stmt, postId, comment, authorID)
	if err != nil {
		return err
	}
	err = CM.UpdateCommentsCount()
	if err != nil {
		return err
	}
	return nil
}

func (CM *CommentModel) UpdateCommentsCount() error {
	stmt := `UPDATE posts
	SET comments_count = (
		SELECT COUNT(*)
		FROM comments
		WHERE comments.postID = posts.postID
	);`

	_, err := CM.DB.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

func (CM *CommentModel) GetComments(user UserModelInterface, postId int, authorID int) ([]*PostComments, error) {
	stmt := `SELECT postID, commentID, comment, authorID,likes_count, dislikes_count FROM comments WHERE postID = ?;`
	rows, err := CM.DB.Query(stmt, postId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []*PostComments

	for rows.Next() {

		c := &PostComments{}

		err = rows.Scan(&c.PostID, &c.CommentID, &c.Comment, &c.AuthorID, &c.Likes, &c.Dislikes)

		if err != nil {
			return nil, err
		}
		c.IsLiked = CM.IsCommentLikedByUser(authorID, c.CommentID)
		c.IsDisliked = CM.IsCommentDislikedByUser(authorID, c.CommentID)

		c.Author, err = user.GetUserNamebyUserID(c.AuthorID)

		if err != nil {
			return nil, err
		}

		comments = append([]*PostComments{c}, comments...)
	}

	if len(comments) == 0 {
		return []*PostComments{}, nil
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (CM *CommentModel) CommentLikeInsert(likeData CommentLikeData, likedBy int) error {
	var commentExists bool
	err := CM.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM comments WHERE commentID = ?)", likeData.CommentID).Scan(&commentExists)
	if err != nil {
		return err
	}

	if !commentExists {
		return errors.New("comment does not exist")
	}

	var count int
	err = CM.DB.QueryRow("SELECT COUNT(*) FROM comment_likes WHERE commentid = ? AND likedby = ?", likeData.CommentID, likedBy).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		_, err := CM.DB.Exec(`DELETE FROM comment_likes WHERE likedby = ? AND commentid = ?`, likedBy, likeData.CommentID)
		if err != nil {
			return err
		}
		return nil
	} else {

		_, err := CM.DB.Exec(`INSERT INTO comment_likes (postid, commentid, likedby) VALUES (?, ?, ?)`, likeData.PostID, likeData.CommentID, likedBy)
		if err != nil {
			return err
		}

		err = CM.RemoveCommentDislike(likeData.CommentID, likedBy)
		if err != nil {
			return err
		}
	}

	return nil
}

func (CM *CommentModel) RemoveCommentLike(commentid int, likedBy int) error {
	if !CM.IsCommentDislikedByUser(likedBy, commentid) {
		return errors.New("comment is already unliked")
	}

	stmt := `DELETE FROM comment_likes WHERE likedby = ? AND commentid = ?`
	_, err := CM.DB.Exec(stmt, likedBy, commentid)
	if err != nil {
		return err
	}

	return nil
}

func (CM *CommentModel) IsCommentLikedByUser(user int, commentid int) bool {
	stmt := `SELECT EXISTS (SELECT * FROM comment_likes WHERE likedby = ? AND commentid = ?)`

	var exists bool

	err := CM.DB.QueryRow(stmt, user, commentid).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (CM *CommentModel) CommentDislikeInsert(dislikeData CommentDislikeData, dislikedBy int) error {
	var commentExists bool
	err := CM.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM comments WHERE commentID = ?)", dislikeData.CommentID).Scan(&commentExists)
	if err != nil {
		return err
	}

	if !commentExists {
		return errors.New("comment does not exist")
	}

	var count int
	err = CM.DB.QueryRow("SELECT COUNT(*) FROM comment_dislikes WHERE commentid = ? AND dislikedby = ?", dislikeData.CommentID, dislikedBy).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		_, err := CM.DB.Exec(`DELETE FROM comment_dislikes WHERE dislikedby = ? AND commentid = ?`, dislikedBy, dislikeData.CommentID)
		if err != nil {
			return err
		}
		return nil
	} else {

		_, err := CM.DB.Exec(`INSERT INTO comment_dislikes (postid, commentid, dislikedby) VALUES (?, ?, ?)`, dislikeData.PostID, dislikeData.CommentID, dislikedBy)
		if err != nil {
			return err
		}

		err = CM.RemoveCommentLike(dislikeData.CommentID, dislikedBy)
		if err != nil {
			return err
		}
	}

	return nil
}

func (CM *CommentModel) RemoveCommentDislike(commentid int, likedBy int) error {
	if !CM.IsCommentDislikedByUser(likedBy, commentid) {
		return errors.New("comment is already undisliked")
	}
	stmt := `DELETE FROM comment_dislikes WHERE dislikedby = ? AND commentid = ?`
	_, err := CM.DB.Exec(stmt, likedBy, commentid)
	if err != nil {
		return err
	}
	return nil
}

func (CM *CommentModel) IsCommentDislikedByUser(userID int, commentid int) bool {
	stmt := `SELECT EXISTS (SELECT * FROM comment_dislikes WHERE dislikedby = ? AND commentid = ?)`

	var exists bool

	err := CM.DB.QueryRow(stmt, userID, commentid).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (CM *CommentModel) UpdateCommentLikesCount() error {
	stmt := `UPDATE comments
	SET likes_count = (
		SELECT COUNT(*)
		FROM comment_likes
		WHERE comment_likes.commentid = comments.commentID
	), 
	dislikes_count = (
		SELECT COUNT(*)
		FROM comment_dislikes
		WHERE comment_dislikes.commentid = comments.commentID
	);`
	_, err := CM.DB.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}
