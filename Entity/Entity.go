package entity

import (
	_ "github.com/jinzhu/gorm"
)

type Employee struct {
	Name       string   `gorm:"column:name;type:varchar;not null"`
	Age        int64    `gorm:"column:age;type:int4;not null"`
	Location   Location `gorm:"foreignKey:location_id;reference:location_id"`
	Gender     string   `gorm:"column:gender;type:varchar;not null"`
	ID         string   `gorm:"column:id;primaryKey;type:uuid;not null"`
	CompanyID  string   `gorm:"column:company_id;type:uuid;not null"`
	LocationID string   `gorm:"column:location_id;type:uuid;not null"`
}

type Company struct {
	ListOfEmployees []Employee `gorm:"foreignKey:company_id"`
	Location        Location   `gorm:"foreignKey:location_id;reference:location_id"`
	LocationID      string
	Name            string `gorm:"column:name;type:varchar;not null"`
	ID              string `gorm:"column:id;primaryKey;type:uuid;not null"`
}

type Location struct {
	Longitude int    `gorm:"column:longitude;type:int4"`
	Latitude  int    `gorm:"column:latitude;type:int4"`
	City      string `gorm:"column:city;type:varchar"`
	ID        string `gorm:"column:id;primaryKey;type:uuid;not null"`
}
