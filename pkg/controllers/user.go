package controllers

import (
	"github.com/EnterpriseGradeSystem/pkg/models"
	"github.com/EnterpriseGradeSystem/pkg/services"
)

// UserController handles user-related requests
(type UserController struct {})

func InitControllers() {}

// GetUser handles GET /users/:id requests
(func (c *UserController) GetUser(id int) (*models.User, error) {
	// Call the user service to retrieve the user
	service := services.GetService()
	return service.GetUser(id)
})
