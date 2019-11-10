package main

import (
	"fmt"
	"github.com/api-metegol/models"
	"github.com/api-metegol/routers"
	"github.com/api-metegol/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

func obtainDbConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/metegol_db?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Participant{}, &models.Tournament{}, &models.Rule{}, &models.Participation{})

	db.Model(&models.Participation{}).AddForeignKey("participant_id", "participants(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Participation{}).AddForeignKey("tournament_id", "tournaments(id)", "RESTRICT", "RESTRICT")

	return db
}

func main() {
	router := gin.Default()

	checkEnvironmentVariables()

	db := obtainDbConnection()

	deps := utils.Dependencies{Db: db}

	routers.InitializeRoutes(router, deps)

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT"))); err != nil {
		panic(err)
	}
}
