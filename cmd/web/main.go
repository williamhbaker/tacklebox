package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/namsral/flag"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var port int
	flag.IntVar(&port, "port", 3000, "Port to start the server listening on")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     fmt.Sprintf(":%d", port),
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %d\n", port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
