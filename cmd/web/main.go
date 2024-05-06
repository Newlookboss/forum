package main

import (
	"flag"
	"log"
	"mkassymk/forum/internal/models"
	"net/http"
	"os"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	users         models.UserModelInterface
	posts         models.PostModelInterface
	likes         models.LikeModelInterface
	sessions      models.SessionModelInterface
	comments      models.CommentModelInterface
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", "4000", "HTTP network address")
	dsn := flag.String("dsn", "forum.db", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "\033[92mINFO\033[0m\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "\033[91mERROR\033[0m\t", log.Ldate|log.Ltime|log.Lshortfile)

	// hashedError, err := bcrypt.GenerateFromPassword([]byte("user/login"), 12)
	// fmt.Println(string(hashedError))
	// hashedError1, err := bcrypt.GenerateFromPassword([]byte("user/signup"), 12)
	// fmt.Println(string(hashedError1))

	db, err := models.CreateDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		users:    &models.UserModel{DB: db},
		posts:    &models.PostModel{DB: db},
		likes:    &models.LikeModel{DB: db},
		sessions: &models.SessionModel{DB: db},
		comments: &models.CommentModel{DB: db},

		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     ":" + *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on http://localhost:%s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
