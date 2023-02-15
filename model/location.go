package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID     string  `json:"id" gorm:"type:varchar(255);primary_key"`
	Ent    int     `json:"cve_ent"`
	Nombre string  `json:"name"`
	Lat    float32 `json:"latitude"`
	Lon    float32 `json:"longitude"`
}

type LocationPayloadDto struct {
	Ent    int     `json:"cve_ent" binding:"required"`
	Nombre string  `json:"name" binding:"required"`
	Lat    float32 `json:"latitude" binding:"required"`
	Lon    float32 `json:"longitude" binding:"required"`
}

type LocationUpdateDto struct {
	Ent    int     `json:"cve_ent"`
	Nombre string  `json:"name"`
	Lat    float32 `json:"latitude"`
	Lon    float32 `json:"longitude"`
}

func (location *Location) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	location.ID = uuid.String()
	return
}
