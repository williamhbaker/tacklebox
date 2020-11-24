package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/namsral/flag"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	col      *mongo.Collection
}

func main() {
	var port int
	flag.IntVar(&port, "port", 3000, "Port to start the server listening on")
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

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		col:      col,
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
