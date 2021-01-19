package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/muhammednagy/pipedirve-challenge/models"
)

type PersonHandler struct {
	config models.Config
	db     *gorm.DB
}

func NewPersonHandler(config models.Config, db *gorm.DB) *PersonHandler {
	return &PersonHandler{config: config, db: db}
}
