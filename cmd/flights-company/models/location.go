package models

// Location ubicacion
type Location struct {
	ID int64 `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`

}

// NewLocation construye un nuevo locacion
func NewLocation(id int64 ,name string) *Location {
	return &Location{ID: id, Name: name}
}

