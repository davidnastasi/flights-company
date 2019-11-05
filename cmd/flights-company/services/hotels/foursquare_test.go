package hotels

import (
	"github.com/davidnastasi/flights-company/cmd/flights-company/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoursquare_GetHotels(t *testing.T){

	config.Config.FoursquareClientId= "HACVIHTUOMFKVK5HWQ0J0JCOKQAA2CSAVFS0LFQVN14EESS2"
	config.Config.FoursquareClientSecret="50ITRVSKRB1GH2YWOBBQWZS5BEDVEIWN3Z2YABJEI454V2JZ"

	hotels , err := GetHotels("Buenos Aires, Argentina")


	assert.Nil(t, err)
	assert.NotEmpty(t, hotels)
}


func TestFoursquare_GetHotelsInvalidLocation(t *testing.T){
	config.Config.FoursquareClientId= "HACVIHTUOMFKVK5HWQ0J0JCOKQAA2CSAVFS0LFQVN14EESS2"
	config.Config.FoursquareClientSecret="50ITRVSKRB1GH2YWOBBQWZS5BEDVEIWN3Z2YABJEI454V2JZ"

	hotels , err := GetHotels("BuenoAires, Argenta")

	assert.NotNil(t, err)
	assert.Empty(t, hotels)
}


