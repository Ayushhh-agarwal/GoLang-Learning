package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/razorpay/GoLang-Learning/first-api/config"
	"github.com/razorpay/GoLang-Learning/first-api/models"
	"github.com/razorpay/GoLang-Learning/first-api/routes"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()

	r.Run()
}
