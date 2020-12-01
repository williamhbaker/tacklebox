package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/golangcollege/sessions"
	"github.com/namsral/flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/wbaker85/tacklebox/pkg/models/mongodb"
	"github.com/wbaker85/tacklebox/pkg/models/postgres"
)

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	session     *sessions.Session
	hooks       *mongodb.HookModel
	hookRecords *postgres.HookRecordModel
	users       *postgres.UserModel
	bins        *postgres.BinModel
}

func main() {
	var port int
	var secret string
	var pgDSN string
	var mongoDSN string

	flag.IntVar(&port, "port", 4000, "Port to start the server listening on")
	flag.StringVar(&secret, "secret", "cookiesecret!", "Secret key for session cookies")
	flag.StringVar(&pgDSN, "pgDSN", "postgres://postgres:postgres@localhost/postgres?sslmode=disable", "Connection string for postgres")
	flag.StringVar(&mongoDSN, "mongoDSN", "mongodb://localhost:27017", "Connection string for MongoDB")
	flag.Parse()

	addr := fmt.Sprintf(":%v", port)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	ctx := context.TODO()
	mongoClient, err := configMongoClient(ctx, mongoDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	col := mongoClient.Database("hooks").Collection("hooks")

	pgDB, err := openPostgres(pgDSN)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer pgDB.Close()

	session := sessions.New([]byte(secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteStrictMode

	app := &application{
		errorLog:    errorLog,
		infoLog:     infoLog,
		session:     session,
		hooks:       &mongodb.HookModel{Col: col, Ctx: &ctx},
		hookRecords: &postgres.HookRecordModel{DB: pgDB},
		users:       &postgres.UserModel{DB: pgDB},
		bins:        &postgres.BinModel{DB: pgDB},
	}

	srv := &http.Server{
		Addr:         addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func configMongoClient(ctx context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func openPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
