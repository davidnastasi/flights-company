package models

// Hotel estructura para los hotel
type Hotel struct{
	Name string `json:"name"`
	Address string `json:"address"`
}

// NewHotel construye un nuevo hotel
func NewHotel(name string, direction string) *Hotel {
	return &Hotel{Name: name, Address: direction}
}
