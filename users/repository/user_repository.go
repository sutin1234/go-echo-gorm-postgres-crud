package repository

import (
	"postgres_api/users/models"

)

// Output struct
type Output struct {
	Result interface{}
	Error  error
}

// UserRepository interface
type UserRepository interface {
	Save(*models.User) Output
	Delete(*models.User) Output
	FindByID(string) Output
	FindAll() Output
}
