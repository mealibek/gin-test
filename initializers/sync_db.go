package initializers

import (
	db "github.com/mealibek/gin-test/db"
)

func SyncDatabase() {
	DB.AutoMigrate(&db.User{})
}
