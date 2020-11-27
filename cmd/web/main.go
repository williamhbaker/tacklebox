package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/caddyserver/certmagic"
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
	var domain string
	var email string
	var staging bool

	flag.IntVar(&port, "port", 443, "Port to start the server listening on")
	flag.StringVar(&secret, "secret", "cookiesecret!", "Secret key for session cookies")
	flag.StringVar(&pgDSN, "pgDSN", "postgres://postgres:postgres@localhost/postgres?sslmode=disable", "Connection string for postgres")
	flag.StringVar(&mongoDSN, "mongoDSN", "mongodb://localhost:27017", "Connection string for MongoDB")
	flag.StringVar(&domain, "domain", "", "Domain to request a certificate for")
	flag.StringVar(&email, "email", "email@domain.com", "Email to use with lets encrypt")
	flag.BoolVar(&staging, "staging", true, "Staging for certs - true for staging, false for real")
	flag.Parse()

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

	app := &application{
		errorLog:    errorLog,
		infoLog:     infoLog,
		session:     session,
		hooks:       &mongodb.HookModel{Col: col, Ctx: &ctx},
		hookRecords: &postgres.HookRecordModel{DB: pgDB},
		users:       &postgres.UserModel{DB: pgDB},
		bins:        &postgres.BinModel{DB: pgDB},
	}

	certmagic.DefaultACME.Agreed = true
	certmagic.DefaultACME.Email = email
	if staging {
		certmagic.DefaultACME.CA = certmagic.LetsEncryptStagingCA
	}

	infoLog.Printf("Server starting")
	err = certmagic.HTTPS([]string{domain}, app.routes())
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
