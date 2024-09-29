package main

import (
	"flag"
	"fmt"
	"log"
	"teststack/api"
	"teststack/database"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "server address")
	flag.Parse()
	database, err := database.NewSqliteDB()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Init()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(*listenAddr, database)
	fmt.Println("Server running on port", *listenAddr)
	log.Fatal(server.Start())
}
