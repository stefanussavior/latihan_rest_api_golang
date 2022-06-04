package model

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Nama          string `json:"nama"`
	Umur          string `json:"umur"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
}
