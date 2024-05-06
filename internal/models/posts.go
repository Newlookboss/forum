package models

import (
	"database/sql"
	"errors"
	"html"
	"strings"
	"time"
)

type Post struct {
	PostID     int
	UserID     int
	Author     string
	Title      string
	Content    string
	Likes      int
	Dislikes   int
	Comments   int
	Tags       string
	IsLiked    bool
	IsDisliked bool
	Created    time.Time
	IsLogged   bool
}

type PostModel struct {
	DB *sql.DB
}

type PostModelInterface interface {
	Insert(userID int, author string, title string, content string, tags []string) (int, error)
	CategoryInsert(postid int64, categories []string) error
	GetPostbyPostID(id int) (*Post, error)
	GetUserPosts(id int) (map[int]*Post, error)
	GetUserLikedPosts(userID int) (map[int]*Post, error)
	Latest(RM LikeModelInterface) (map[int]*Post, error)
	FilteredPosts(categories []string) (map[int]*Post, error)
}

func (PM *PostModel) Insert(userID int, author string, title string, content string, tags []string) (int, error) {
	title = html.EscapeString(title)
	content = html.EscapeString(content)

	stmt := "INSERT INTO posts (userID, author, title, content, created, likes_count, dislikes_count, comments_count, tags)\n\tVALUES(?,?, ?, ?, CURRENT_TIMESTAMP, \"0\", \"0\", \"0\", ?)"

	result, err := PM.DB.Exec(stmt, userID, author, title, content, strings.Join(tags, " "))
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	err = PM.CategoryInsert(id, tags)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (PM *PostModel) CategoryInsert(postid int64, categories []string) error {
	stmt2 := `INSERT INTO categories (postid, category) VALUES (?, ?)`

	for _, category := range categories {
		_, err := PM.DB.Exec(stmt2, postid, category)
		if err != nil {
			return err
		}
	}
	return nil
}

func (PM *PostModel) GetPostbyPostID(id int) (*Post, error) {
	stmt := `SELECT postID, author, title, content, created, likes_count, dislikes_count, comments_count, tags FROM posts
    WHERE postID = ?`

	row := PM.DB.QueryRow(stmt, id)
	s := &Post{}

	err := row.Scan(&s.PostID, &s.Author, &s.Title, &s.Content, &s.Created, &s.Likes, &s.Dislikes, &s.Comments, &s.Tags)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (PM *PostModel) GetUserPosts(id int) (map[int]*Post, error) {
	stmt := `SELECT postID, author, title, content, created, likes_count, dislikes_count, comments_count, tags FROM posts
    WHERE userID = ?`

	rows, err := PM.DB.Query(stmt, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	posts := map[int]*Post{}

	for rows.Next() {
		s := &Post{}

		err = rows.Scan(&s.PostID, &s.Author, &s.Title, &s.Content, &s.Created, &s.Likes, &s.Dislikes, &s.Comments, &s.Tags)
		if err != nil {
			return nil, err
		}
		posts[s.PostID] = s
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (PM *PostModel) GetUserLikedPosts(userID int) (map[int]*Post, error) {
	stmt := `SELECT p.postID, p.author, p.title, p.content, p.created, p.likes_count, p.dislikes_count, p.comments_count, tags 
             FROM posts p
             JOIN likes l ON p.postID = l.postid
             WHERE l.likedby = ?`

	rows, err := PM.DB.Query(stmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	likedPosts := map[int]*Post{}

	for rows.Next() {
		s := &Post{}
		err = rows.Scan(&s.PostID, &s.Author, &s.Title, &s.Content, &s.Created, &s.Likes, &s.Dislikes, &s.Comments, &s.Tags)
		if err != nil {
			return nil, err
		}
		likedPosts[s.PostID] = s

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return likedPosts, nil
}

func (PM *PostModel) Latest(RM LikeModelInterface) (map[int]*Post, error) {
	stmt := `SELECT postID, userID, author, title, content, created, likes_count, dislikes_count, comments_count, tags FROM posts
	WHERE postID <= CURRENT_TIMESTAMP ORDER BY postID DESC`

	rows, err := PM.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := map[int]*Post{}

	for rows.Next() {
		s := &Post{}

		err = rows.Scan(&s.PostID, &s.UserID, &s.Author, &s.Title, &s.Content, &s.Created, &s.Likes, &s.Dislikes, &s.Comments, &s.Tags)
		if err != nil {
			return nil, err
		}
		// to check this function
		posts[s.PostID*(-1)] = s
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (PM *PostModel) FilteredPosts(categories []string) (map[int]*Post, error) {
	posts := map[int]*Post{}
	stmt := `SELECT posts.postID, userID, author, title, content, created, likes_count, dislikes_count, comments_count, tags 
	FROM posts JOIN categories ON posts.postID = categories.postid 
	WHERE categories.category = ?;`

	for _, category := range categories {
		rows, err := PM.DB.Query(stmt, category)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			s := &Post{}

			err = rows.Scan(&s.PostID, &s.UserID, &s.Author, &s.Title, &s.Content, &s.Created, &s.Likes, &s.Dislikes, &s.Comments, &s.Tags)
			if err != nil {
				return nil, err
			}

			posts[s.PostID*(-1)] = s
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	return posts, nil
}
