package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	initJob()
	router := NewRouter()
	port := 8383
	fmt.Printf("Initializing the database\n")
	InitDb()
	defer db.Close()
	logger.Warn("Listening on ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
