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
	flag.IntVar(&port, "port", 3000, "Port to start the server listening on")
	flag.StringVar(&secret, "secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key for session cookies")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	ctx := context.TODO()
	mongoClient, err := configMongoClient(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	col := mongoClient.Database("hooks").Collection("hooks")

	pgDB, err := openPostgres("postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		errorLog.Fatal(err)
	}
	defer pgDB.Close()

	session := sessions.New([]byte(secret))
	session.Lifetime = 12 * time.Hour

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
		Addr:     fmt.Sprintf(":%d", port),
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %d\n", port)
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
