package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/edwlarkey/cardamom/pkg/config"
	"github.com/edwlarkey/cardamom/pkg/db"
	"github.com/edwlarkey/cardamom/pkg/models"
	"github.com/golangcollege/sessions"
)

type contextKey string

var contextKeyUser = contextKey("user")

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	session  *sessions.Session
	config   config.Config
	db       interface {
		Connect(string, string) error
		Migrate() error
		LatestBookmarks() ([]*models.Bookmark, error)
		GetBookmarks(*models.User) ([]*models.Bookmark, error)
		GetBookmark(uint, *models.User) (*models.Bookmark, error)
		InsertBookmark(*models.Bookmark) error
		UpdateBookmark(*models.Bookmark) error
		GetTags(*models.User) ([]*models.Tag, error)
		GetTagByName(string, *models.User) (*models.Tag, error)
		CreateIfNotExists(string) (*models.Tag, error)
		InsertTag(*models.Tag) (uint, error)
		InsertUser(string, string, string) error
		AuthenticateUser(string, string) (*models.User, error)
		GetUser(uint) (*models.User, error)
	}
	templates *templates
}

func main() {
	configFile := flag.String("config", "./config.toml", "Config file")
	flag.Parse()

	var conf config.Config
	if _, err := toml.DecodeFile(*configFile, &conf); err != nil {
		fmt.Println(err)
	}

	var dsn string
	switch conf.Database.Dialect {
	case "sqlite":
		dsn = conf.Database.Database
	default:
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.Database.Server, conf.Database.Port, conf.Database.User, conf.Database.Database, conf.Database.Password)
	}

	addr := fmt.Sprintf("%s:%s", conf.Web.IP, conf.Web.Port)

	infoLog := log.New(os.Stdout, "INFO\t", log.LUTC|log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new session manager, passing in the secret key as the parameter.
	session := sessions.New([]byte(conf.Web.Secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	t, err := initTemplates()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:  errorLog,
		infoLog:   infoLog,
		session:   session,
		config:    conf,
		templates: t,
		db:        &db.DB{},
	}

	// Connect to the DB
	err = app.db.Connect(conf.Database.Dialect, dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	err = app.db.Migrate()
	if err != nil {
		errorLog.Fatal(err)
	}

	// Set up http server, including app routes
	srv := &http.Server{
		Addr:         addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		infoLog.Printf("Starting server on %s", addr)
		if err := srv.ListenAndServe(); err != nil {
			errorLog.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err = srv.Shutdown(ctx)
	if err != nil {
		errorLog.Println(err)
	}
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	infoLog.Println("Shutting down. Bye!")
	os.Exit(0)
}
