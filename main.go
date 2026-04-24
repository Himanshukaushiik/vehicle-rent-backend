package main

import (
	"awesomeProject1/config"
	"awesomeProject1/internal/rent"
	"awesomeProject1/internal/user"
	"awesomeProject1/internal/vehicles"
	"awesomeProject1/models"
	"awesomeProject1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()

	vc := vehicles.InitializeVehicleController()
	rc := rent.InitializeRentController()
	uc := user.InitializeUserController()

	routes.SetupRoutes(r, vc, rc, uc)

	config.DB.AutoMigrate(
		&models.Vehicle{},
		&models.User{},
		&models.Rent{},
	)

	r.Run(":8080")
}
