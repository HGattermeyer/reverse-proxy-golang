package data

import (
	"proxy-golang/internal/models"

	"gorm.io/gorm"
)

func GetServerByUri(uri string, db *gorm.DB) (models.Server, error) {
	var server models.Server

	err := db.Model(&models.Server{}).Preload("RedirectServer").Where("uri = ?", uri).Find(&server).Error
	return server, err
}

func GetAllServers(db *gorm.DB) ([]models.Server, error) {
	var servers []models.Server

	err := db.Model(&models.Server{}).Preload("RedirectServer").Find(&servers).Error

	return servers, err
}

func SaveOrUpdateServer(server models.Server, db *gorm.DB) models.Server {

	db.Save(&server)

	return server
}

func DeleteServerByUri(uri string, db *gorm.DB) error {
	server, err := GetServerByUri(uri, db)

	// Delete associated RedirectServers
	db.Delete(&server.RedirectServer)
	db.Delete(&server)

	return err
}
