package config

import "github.com/jinzhu/gorm"

// Config application configurations
var Config appConfig

type appConfig struct {
	// the shared DB ORM object
	DB *gorm.DB
	// the error thrown be GORM when using DB ORM object
	DBErr error
	// the data source name (DSN) for connecting to the database. required.
	DSN string
	// the server port. Defaults to 8080
	WebServer string
	// Foursquare client id
	FoursquareClientId string
	// Foursquare client secreat
	FoursquareClientSecret string
	// Endpoint reserves
	ReservationsEndpoint string

}
