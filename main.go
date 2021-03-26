package main

import (
	"flag"
	"log"
	"net/http"

	"os"
)

type HandlerStrust struct {
	printError *log.Logger
	printInfo  *log.Logger
	views      *Views
}

func main() {

	networkAddress := flag.String("networkAddress", ":8080", "HTTP network address")

	flag.Parse()

	printInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	printError := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := initdb("postgres://postgres:your_password@localhost/customersdb?sslmode=disable")

	if err != nil {
		printError.Fatal(err)
	}

	defer db.Close()

	h := &HandlerStrust{
		printError: printError,
		printInfo:  printInfo,
		views:      &Views{dtbs: db},
	}

	ConnectToServer := &http.Server{
		Addr:     *networkAddress,
		ErrorLog: printError,
		Handler:  h.Router(),
	}
	printInfo.Printf("Starting server on ................. %s", *networkAddress)
	printInfo.Printf("Connected to database successfully!")
	err = ConnectToServer.ListenAndServe()
	printError.Fatal(err)
}
