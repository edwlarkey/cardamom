package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"git.sr.ht/~edwlarkey/cardamom/internal/models/sqlite"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "modernc.org/sqlite"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	bookmarks      *sqlite.BookmarkModel
	users          *sqlite.UserModel
	tags           *sqlite.TagModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "cardamom.sqlite?_time_format=sqlite", "data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dbx, err := sqlite.OpenDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer dbx.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Store = sqlite3store.New(dbx.DB)
	sessionManager.Lifetime = 12 * time.Hour

	sessionManager.Cookie.Secure = true

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		bookmarks:      &sqlite.BookmarkModel{DB: dbx},
		users:          &sqlite.UserModel{DB: dbx},
		tags:           &sqlite.TagModel{DB: dbx},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
