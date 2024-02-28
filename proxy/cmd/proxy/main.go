package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/masharpik/HTTPproxy/internal/provider"
	"github.com/masharpik/HTTPproxy/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		err = fmt.Errorf("godotenv.Load: %w", err)
		log.Fatalln(err.Error())
	}

	handler := provider.InitHandler()

	if err = server.Init(os.Getenv("PROXY_HOST"), os.Getenv("PROXY_PORT"), handler); err != nil {
		err = fmt.Errorf("server.Init: %w", err)
		log.Fatalln(err.Error())
	}
}
