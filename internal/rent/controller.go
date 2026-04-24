package rent

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RentController struct {
	service domain.RentService
}

func NewRentController(s domain.RentService) *RentController {
	return &RentController{service: s}
}

func (rc *RentController) CreateRent(c *gin.Context) {
	var rent models.Rent

	if err := c.ShouldBindJSON(&rent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := rc.service.CreateRent(&rent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rent)
}

func (rc *RentController) GetAllRents(c *gin.Context) {
	rents, err := rc.service.GetAllRent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": rents,
	})
}
