package models

import "gorm.io/gorm"

type ServerFile struct {
	Uri             string
	RedirectCount   int
	RedirectServers []string
}

type Server struct {
	gorm.Model
	Uri            string
	RedirectCount  int
	RedirectServer []RedirectServer
}

// type Server struct {
// 	Uri             string   `json:"uri"`
// 	RedirectCount   int      `json:"redirectCount"`
// 	RedirectServers []string `json:"redirectServers"`
// }
