package main

import (
	"github.com/asaskevince/govalidator"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string
	Email      string
	CustomerID string
}

