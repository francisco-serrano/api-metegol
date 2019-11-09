package main

import (
	"fmt"
	"github.com/api-metegol/routers"
	"github.com/gin-gonic/gin"
	"os"
)

func checkEnvironmentVariables() {
	envVars := []string{
		"API_PORT",
	}

	for _, v := range envVars {
		if myVar := os.Getenv(v); myVar == "" {
			panic(fmt.Sprintf("%s not provided", v))
		}
	}
}

func main() {
	checkEnvironmentVariables()

	router := gin.Default()

	routers.InitializeRoutes(router)

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT"))); err != nil {
		panic(err)
	}
}
