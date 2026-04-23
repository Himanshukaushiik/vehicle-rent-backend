package rent

import (
	"awesomeProject1/models"

	"github.com/gin-gonic/gin"
)

type RentController struct {
	service *RentService
}

func NewRentController(s *RentService) *RentController {
	return &RentController{service: s}
}

func (rc *RentController) CreateRent(c *gin.Context) {
	var rent models.Rent

	if err := c.ShouldBindJSON(&rent); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	if err := rc.service.CreateRent(&rent); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, rent)
}

func (rc *RentController) GetAllRent(c *gin.Context) {
	var rents []models.Rent
	rents, err := rc.service.GetAllRent()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": rents,
	})
}
