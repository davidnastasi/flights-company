package main

import (
	"github.com/davidnastasi/flights-company/cmd/flights-company/api"
	"github.com/davidnastasi/flights-company/cmd/flights-company/config"
	"github.com/davidnastasi/flights-company/cmd/flights-company/services/reservation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

const CLIENT_ID  = "HACVIHTUOMFKVK5HWQ0J0JCOKQAA2CSAVFS0LFQVN14EESS2"
const CLIENT_SECRET = "50ITRVSKRB1GH2YWOBBQWZS5BEDVEIWN3Z2YABJEI454V2JZ"


func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("api/v1")
	{
		v1.GET("/destinations", api.GetDestinations)
	}


	loadConfig()

	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}


	cronJob := cron.New()

	_, err := cronJob.AddFunc("@every 5s", func(){
		if err := reservation.GetAll() ; err != nil {
			log.Println(err)
		}
	})

	if err != nil {
		panic(err)
	}
	cronJob.Start()

	defer cronJob.Stop()
	defer config.Config.DB.Close()

	err = r.Run(config.Config.WebServer)
	if err != nil {
		panic(err)
	}else{
		log.Println("Server started successfully")
	}
}

func loadConfig() {
	config.Config.WebServer = getEnv("APP_SERVER_URL", ":8080")
	config.Config.DSN = getEnv("APP_DB_DSN", "host=localhost port=5432 user=postgres dbname=flights password=s3cret sslmode=disable" )
	config.Config.FoursquareClientId= getEnv("APP_CLIENT_ID", CLIENT_ID)
	config.Config.FoursquareClientSecret= getEnv("APP_CLIENT_SECRET", CLIENT_SECRET)
	config.Config.ReservationsEndpoint = getEnv("APP_RESERVATION_ENDPOINT", "https://brubank-flights.herokuapp.com/flight-reservations")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}