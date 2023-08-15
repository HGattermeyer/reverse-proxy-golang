package models

import "gorm.io/gorm"

type RedirectServer struct {
	gorm.Model
	ServerID uint
	Server   string
}
