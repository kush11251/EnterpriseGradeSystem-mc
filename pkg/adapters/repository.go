package adapters

import (
	"github.com/EnterpriseGradeSystem/pkg/models"
)

// UserRepository is an adapter for user storage
(type UserRepository interface {
	GetUser(id int) (*models.User, error)
})
