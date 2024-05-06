package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"mkassymk/forum/internal/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// status := r.Context().Value("data").(*contextData)
	if (r.Method != http.MethodGet) && (r.Method != http.MethodPost) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.URL.RawQuery != "" {
		if r.URL.RawQuery != "r="+models.JSSuccesfulSignup && r.URL.RawQuery != "r="+models.JSDuplicateEmail && r.URL.RawQuery != "r="+models.JSLogin && r.URL.RawQuery != "r="+models.JSSignup && r.URL.RawQuery != "r="+models.JSInvalidCredentials {
			app.notFound(w)
			return
		}
	}

	posts, err := app.posts.Latest(app.likes)
	if err != nil {
		app.serverError(w, err)
		return
	}

	switch r.Method {
	case http.MethodGet:
		if r.URL.Path != "/" {
			app.notFound(w)
			return
		}
	case http.MethodPost:
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				app.clientError(w, http.StatusBadRequest)
				return
			}
			categories := r.Form["category"]
			posts, err = app.FilteredPosts(categories)
			if err != nil {
				app.serverError(w, err)
				return
			}
		}

	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	data := app.newTemplateData(r)
	data.Posts = posts
	if data.IsLogged {
		for i, j := range data.Posts {
			data.Posts[i].IsLiked = app.likes.IsLikedByUser(data.User.UserID, j.PostID)
			data.Posts[i].IsDisliked = app.likes.IsDislikedByUser(data.User.UserID, j.PostID)
			data.Posts[i].IsLogged = true
		}
	}
	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) FilteredPosts(categories []string) (map[int]*models.Post, error) {
	if len(categories) == 0 {
		posts, err := app.posts.Latest(app.likes)
		if err != nil {
			return nil, err
		}
		return posts, nil
	}
	posts, err := app.posts.FilteredPosts(categories)
	if err != nil {
		return nil, err
	}

	// Iterate through each post and check if all categories are present in the tags.
	filteredPosts := make(map[int]*models.Post)
	for key, post := range posts {
		postCategories := strings.Split(post.Tags, " ")

		if containsAllCategories(postCategories, categories) {
			filteredPosts[key] = post
		}
	}

	return filteredPosts, nil
}

func containsAllCategories(tags []string, categories []string) bool {
	categorySet := make(map[string]struct{})

	// Create a set of categories for quick look-up.
	for _, category := range tags {
		categorySet[category] = struct{}{}
	}
	// Check if all categories are in the tag set.
	for _, tag := range categories {
		if _, ok := categorySet[tag]; !ok {
			return false
		}
	}

	return true
}

// user's handlers

// user signup

type UserSignupForm struct {
	firstname string
	lastname  string
	email     string
	password  string
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.Redirect(w, r, fmt.Sprintf("/?r=%s", models.JSSignup), http.StatusSeeOther)
		return
	case http.MethodPost:

		form := UserSignupForm{
			firstname: r.FormValue("firstname"),
			lastname:  r.FormValue("lastname"),
			email:     strings.ToLower(r.FormValue("email")),
			password:  r.FormValue("password"),
		}

		if form.firstname == "" || form.lastname == "" || form.email == "" || form.password == "" {
			http.Redirect(w, r, "/user/signup", http.StatusSeeOther)
			return
		}

		if len(form.password) < 8 || len(form.password) > 15 {
			http.Redirect(w, r, "/user/signup", http.StatusSeeOther)
			return
		}

		if !strings.Contains(form.email, "@") || !strings.Contains(form.email, ".") {
			http.Redirect(w, r, "/user/signup", http.StatusSeeOther)
			return
		}

		err := app.users.InsertUser(form.firstname, form.lastname, form.email, form.password)
		if err != nil {

			if err.Error() == models.ErrDuplicateEmail.Error() {
				http.Redirect(w, r, fmt.Sprintf("/?r=%s", models.JSDuplicateEmail), http.StatusSeeOther)
			} else {
				app.serverError(w, err)
			}

			return
		}

		http.Redirect(w, r, fmt.Sprintf("/?r=%s", models.JSSuccesfulSignup), http.StatusSeeOther)
		return
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
}

// user login

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.Redirect(w, r, fmt.Sprintf("/?r=%s", models.JSLogin), http.StatusSeeOther)
		return
	case http.MethodPost:
		form := UserSignupForm{
			email:    r.FormValue("email"),
			password: r.FormValue("password"),
		}

		id, err := app.users.Authenticate(form.email, form.password)
		if err != nil {
			if errors.Is(err, models.ErrInvalidCredentials) {
				http.Redirect(w, r, fmt.Sprintf("/?r=%s", models.JSInvalidCredentials), http.StatusSeeOther)
			} else {
				app.serverError(w, err)
			}
			return
		}

		token, err := uuid.NewV4()

		cookie := &http.Cookie{
			Name:  "session",
			Value: token.String(),
			Path:  "/",
		}

		http.SetCookie(w, cookie)

		err = app.sessions.AddToken(id, cookie.Value)
		if err != nil {
			app.serverError(w, err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	return
}

// user logout

func (app *application) userLogout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.clientError(w, http.StatusMethodNotAllowed)
	}
	cookie, err := r.Cookie("session")
	if cookie != nil {
		if err != nil {
			app.serverError(w, err)
			return
		}
		app.sessions.RemoveToken(cookie.Value)
		models.DeleteCookie(w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// user profile

func (app *application) userProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	data := app.newTemplateData(r)
	posts, err := app.posts.GetUserPosts(data.User.UserID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data.Posts = posts
	if data.IsLogged {
		for i, j := range data.Posts {

			data.Posts[i].IsLiked = app.likes.IsLikedByUser(data.User.UserID, j.PostID)
			data.Posts[i].IsDisliked = app.likes.IsDislikedByUser(data.User.UserID, j.PostID)
			j.IsLogged = true
		}
	}
	app.render(w, http.StatusOK, "profile.html", data)
}

func (app *application) userFavourites(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	data := app.newTemplateData(r)
	posts, err := app.posts.GetUserLikedPosts(data.User.UserID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data.Posts = posts
	if data.IsLogged {
		for i, j := range data.Posts {

			data.Posts[i].IsLiked = app.likes.IsLikedByUser(data.User.UserID, j.PostID)
			data.Posts[i].IsDisliked = app.likes.IsDislikedByUser(data.User.UserID, j.PostID)
			j.IsLogged = true
		}
	}
	app.render(w, http.StatusOK, "favs.html", data)
}

// post's handlers

// Post View

func (app *application) postView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	url := r.URL
	// Получаем Query параметры из URL
	query := url.Query()
	// Создаем пустую карту для хранения параметров
	params := make(map[string]string)

	// Извлекаем параметры из Query и добавляем в карту
	for key, values := range query {
		if len(values) > 0 {
			params[key] = values[0] // Берем только первое значение
		}
	}

	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		// Если возникла ошибка при преобразовании или id меньше 1,
		// возвращаем ошибку "Not Found"
		app.notFound(w)
		return
	}

	post, err := app.posts.GetPostbyPostID(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	comments, err := app.comments.GetComments(app.users, id, data.User.UserID)
	if err != nil {

		app.serverError(w, err)
		return
	}
	data.Post = post
	data.Comments = comments

	data.Post.IsLiked = app.likes.IsLikedByUser(data.User.UserID, id)
	data.Post.IsDisliked = app.likes.IsDislikedByUser(data.User.UserID, id)
	app.render(w, http.StatusOK, "view.html", data)
}

// Post Create

type postCreateForm struct {
	Title   string
	Content string
	Tags    []string
}

func (app *application) postCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := postCreateForm{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("description"),
		Tags:    r.Form["category"],
	}

	form.Title = strings.TrimSpace(form.Title)
	form.Content = strings.TrimSpace(form.Content)

	if form.Title == "" || form.Content == "" || len(form.Tags) == 0 {
		app.clientError(w, http.StatusBadRequest)

		return
	}

	user, _ := app.users.GetUserInfo(r)
	_, err = app.posts.Insert(user.UserID, user.Firstname+" "+user.Lastname, form.Title, form.Content, form.Tags)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// post Like

func (app *application) postLike(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if r.URL.Path == "/post/like" {

			user, _ := app.users.GetUserInfo(r)

			var likeData models.UserLikeData

			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&likeData); err != nil {
				app.clientError(w, http.StatusBadRequest)
				return
			}
			err := app.likes.LikeInsert(likeData, user.UserID)
			if err != nil {
				app.serverError(w, err)
				return
			}
			err = app.likes.UpdateLikesCount()
			if err != nil {
				app.serverError(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
	}
}

// post dislike

func (app *application) postDislike(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if r.URL.Path == "/post/dislike" {

			user, _ := app.users.GetUserInfo(r)

			var dislikeData models.UserDislikeData

			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&dislikeData); err != nil {
				app.clientError(w, http.StatusBadRequest)
				return
			}
			err := app.likes.DislikeInsert(dislikeData, user.UserID)
			if err != nil {
				app.serverError(w, err)
				return
			}
			err = app.likes.UpdateLikesCount()
			if err != nil {
				app.serverError(w, err)
				return
			}

			w.WriteHeader(http.StatusOK)
		}
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
	}
}

// post comment
func (app *application) postComment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	user, err := app.users.GetUserInfo(r)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	comment := r.PostForm.Get("comment")
	comment = html.EscapeString(comment)

	trimmedComment := strings.TrimSpace(comment)

	if trimmedComment == "" {
		http.Redirect(w, r, fmt.Sprintf("/post/view?id=%d", id), http.StatusSeeOther)

		return
	}

	err = app.comments.CommentInsert(comment, user.UserID, id)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/post/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) commentLike(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if r.URL.Path == "/post/commentLike" {

			user, err := app.users.GetUserInfo(r)
			if err != nil {
				app.serverError(w, err)
				return
			}

			var commentLikeData models.CommentLikeData

			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&commentLikeData); err != nil {
				app.clientError(w, http.StatusBadRequest)
				return
			}

			err = app.comments.CommentLikeInsert(commentLikeData, user.UserID)
			if err != nil {
				app.serverError(w, err)
				return
			}
			err = app.comments.UpdateCommentLikesCount()
			if err != nil {
				app.serverError(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) commentDislike(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if r.URL.Path == "/post/commentDislike" {

			user, _ := app.users.GetUserInfo(r)

			var commentDislikeData models.CommentDislikeData

			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&commentDislikeData); err != nil {
				app.clientError(w, http.StatusBadRequest)
				return
			}

			err := app.comments.CommentDislikeInsert(commentDislikeData, user.UserID)
			if err != nil {
				app.serverError(w, err)
				return
			}
			err = app.comments.UpdateCommentLikesCount()
			if err != nil {
				app.serverError(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	default:
		app.clientError(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) neuter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			app.notFound(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
