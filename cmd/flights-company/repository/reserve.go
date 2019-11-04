package repository

import (
	"github.com/davidnastasi/flights-company/cmd/flights-company/config"
	"github.com/davidnastasi/flights-company/cmd/flights-company/models"
	"log"
)

// Get obtiene todas las reservas
func GetAll(destination string) ([]*models.Reserve, error) {
	var reserves []*models.Reserve

	err := config.Config.DB.Where("destination = ?", destination).
		Order("date asc").
		Find(&reserves).
		Error
	return reserves, err
}

// Save salva las reservas
func Save(reserves []*models.Reserve) error {
	var err error
	//log.Println(len(reserves))
	//log.Println("*******************************************************************************")
	for _, r := range reserves  {
		log.Println(r)
		config.Config.DB.NewRecord(&r)
		if res := config.Config.DB.Create(r) ; res.Error != nil {
			err = res.Error

		}
	}
	return err
}

