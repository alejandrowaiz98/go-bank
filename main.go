package main

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}

func main() {

	server := NewAPIServer(os.Getenv("server_port"))

	server.Run()

}
