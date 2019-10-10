package main

import (
	"fmt"
	"github.com/api-metegol/routers"
	"os"
)

func main() {
	router := routers.InitializeRouter()

	port := os.Getenv("API_PORT")
	if port == "" {
		panic("port not provided")
	}

	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
