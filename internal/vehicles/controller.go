package vehicles

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	service domain.VehicleService
}

func NewVehicleController(s domain.VehicleService) *VehicleController {
	return &VehicleController{
		service: s,
	}
}

// CREATE

func (vc *VehicleController) CreateVehicle(c *gin.Context) {
	var v models.Vehicle

	if err := c.ShouldBindJSON(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := vc.service.CreateVehicle(&v); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "vehicle created",
		"data":    v,
	})
}

// GET ALL

func (vc *VehicleController) GetVehicles(c *gin.Context) {
	vehicles, err := vc.service.GetVehicles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": vehicles,
	})
}

// GET BY ID
func (vc *VehicleController) GetVehicleById(c *gin.Context) {
	id := c.Param("id")

	vehicle, err := vc.service.GetVehicleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "vehicle not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": vehicle,
	})
}

// DELETE
func (vc *VehicleController) DeleteVehicle(c *gin.Context) {
	id := c.Param("id")

	vehicle, err := vc.service.GetVehicleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "vehicle not found"})
		return
	}

	if err := vc.service.DeleteVehicle(&vehicle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "vehicle deleted",
	})
}
