package main

import (
	"github.com/davidnastasi/flights-company/cmd/flights-company/api"
	"github.com/davidnastasi/flights-company/cmd/flights-company/config"
	"github.com/davidnastasi/flights-company/cmd/flights-company/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("api/v1")
	{
		v1.GET("/destinations", api.GetDestinations)
	}

	//config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=flights password=s3cret sslmode=disable")
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	cronJob := cron.New()

	_, err := cronJob.AddFunc("@every 10m", func(){
		if err := services.GetAll() ; err != nil {
			log.Println(err)
		}
	})

	if err != nil {
		panic(err)
	}
	cronJob.Start()

	defer cronJob.Stop()
	defer config.Config.DB.Close()

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}else{
		log.Println("Server started successfully")
	}
}

