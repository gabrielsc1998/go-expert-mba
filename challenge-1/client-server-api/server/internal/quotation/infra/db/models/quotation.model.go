package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuotationModel struct {
	Id         string `gorm:"primaryKey"`
	Code       string `gorm:"string"`
	Codein     string `gorm:"string"`
	Name       string `gorm:"string"`
	High       string `gorm:"string"`
	Low        string `gorm:"string"`
	VarBid     string `gorm:"string"`
	PctChange  string `gorm:"string"`
	Bid        string `gorm:"string"`
	Ask        string `gorm:"string"`
	Timestamp  string `gorm:"string"`
	CreateDate string `gorm:"string"`
}

func (user *QuotationModel) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = uuid.NewString()
	return
}
