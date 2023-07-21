package initializers

import "example/m/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}