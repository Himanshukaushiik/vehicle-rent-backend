package routes

import (
	"awesomeProject1/internal/rent"
	"awesomeProject1/internal/user"
	"awesomeProject1/internal/vehicles"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, vc *vehicles.VehicleController, rc *rent.RentController, uc *user.UserController) {

	vehiclesGroup := r.Group("/vehicles")
	rentGroup := r.Group("/rent")
	usersGroup := r.Group("/users")

	{
		vehiclesGroup.POST("/", vc.CreateVehicle)
		vehiclesGroup.GET("/", vc.GetVehicles)
		vehiclesGroup.GET("/:id", vc.GetVehicleById)
		vehiclesGroup.DELETE("/:id", vc.DeleteVehicle)

		rentGroup.POST("/", rc.CreateRent)
		rentGroup.GET("/", rc.GetAllRent)

		usersGroup.POST("/", uc.CreateUser)
		usersGroup.GET("/", uc.GetAllUsers)
		usersGroup.DELETE("/:id", uc.DeleteUser)
		usersGroup.PUT("/", uc.UpdateUser)

	}
}
