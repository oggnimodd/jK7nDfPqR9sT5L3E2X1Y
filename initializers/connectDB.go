package initializers

import (
	"fmt"
	"log"

	"github.com/oggnimodd/jK7nDfPqR9sT5L3E2X1Y/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")

	EnableUUID(DB)

	err2 := DB.AutoMigrate(&models.User{}, &models.Post{})
	if err2 != nil {
		fmt.Println("something wrong when trying to migrate user model.")
	}
	fmt.Println("? Migration complete")
}

func EnableUUID(db *gorm.DB) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	fmt.Println("? Enabled uuid-ossp extension")
}
