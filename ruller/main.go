package main

import (
	"fmt"
	"github.com/AlexanderZh/gosurirule/go/database"
	"github.com/AlexanderZh/gosurirule/go/handlers"
	"log"
	"net/http"
	"os"
)

var err error

//
//be sure you provide
func main() {
	fmt.Println("Suricata IDS/IPS rules service")

	log.Printf("Server started")

	router := handlers.NewRouter()
	serve, ok := os.LookupEnv("SERVE")
	if !ok {
		serve = ":9090"
	}

	database.InitDB()

	log.Fatal(http.ListenAndServe(serve, router))
}
