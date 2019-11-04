package api

import (
	"github.com/davidnastasi/flights-company/cmd/flights-company/services"
	"github.com/gin-gonic/gin"
	"net/http"
)


var service = services.NewDestinationService()

// GetDestinations get reserves
func GetDestinations(c *gin.Context){
	//var data []models.Destination
	destination, boolean := c.GetQuery("search")
	if boolean {
		if destination == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, apiError{Code:http.StatusBadRequest, Message:"search query param value must not be empty"})
		} else {
			data, err := service.Get(destination)
			if err != nil {
				if data.Reserves == nil {
					c.AbortWithStatusJSON(http.StatusNotFound, apiError{Code:http.StatusNotFound, Message:err.Error()})
				}else {
					c.AbortWithStatusJSON(http.StatusInternalServerError, apiError{Code:http.StatusInternalServerError, Message:err.Error()})
				}
			}else {
				c.JSON(http.StatusOK,data)
			}

		}
	}else{
		c.AbortWithStatusJSON(http.StatusBadRequest, apiError{Code:http.StatusBadRequest, Message:"search query param not found"})
	}
}

