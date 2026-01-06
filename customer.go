package main

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~Name is required"`
	Email      string `valid:"email"`
	CustomerID string `valid:"matches(^[LMH]\\d{7}$)"`
}
