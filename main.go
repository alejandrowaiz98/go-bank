package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}

func main() {

	postgreStore, err := NewPostgreStore()

	if err != nil {
		panic(err)
	}

	log.Printf("%+v\n", postgreStore)

	//server := NewAPIServer(os.Getenv("server_port"), postgreStore)

	//server.Run()

}
