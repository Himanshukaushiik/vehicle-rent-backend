package user

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service domain.UserService
}

func NewUserController(s domain.UserService) *UserController {
	return &UserController{service: s}
}

// CREATE
func (uc *UserController) CreateUser(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := uc.service.CreateUser(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

// GET ALL
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// DELETE
func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	user.ID = parseID(id)

	if err := uc.service.DeleteUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

// UPDATE
func (uc *UserController) UpdateUser(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := uc.service.UpdateUser(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}
