package services

import (
	"database/sql"
	"github.com/EnterpriseGradeSystem/pkg/config"
	"github.com/EnterpriseGradeSystem/pkg/models"
)

// UserService handles user-related business logic
(type UserService struct {
	DB *sql.DB
})

func GetService() *UserService {
	service := &UserService{}
	// Initialize the database connection
	config := config.InitConfig()
	service.DB, _ = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=example sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password))
	return service
}

// GetUser retrieves a user by ID
(func (s *UserService) GetUser(id int) (*models.User, error) {
	stmt, err := s.DB.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var user models.User
	if err := stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
})
