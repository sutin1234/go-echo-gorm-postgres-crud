package repository

import (
	"postgres_api/users/models"

	"github.com/jinzhu/gorm"

)

// UserRepositoryGorm struct
type UserRepositoryGorm struct {
	db *gorm.DB
}

// NewUserRepositoryGorm func
func NewUserRepositoryGorm(db *gorm.DB) *UserRepositoryGorm {
	return &UserRepositoryGorm{db: db}
}

// Save func
func (u *UserRepositoryGorm) Save(user *models.User) Output {
	err := u.db.Save(user).Error
	if err != nil {
		return Output{Error: err}
	}

	return Output{Result: user}
}

// Delete func
func (u *UserRepositoryGorm) Delete(user *models.User) Output {
	err := u.db.Delete(user).Error
	if err != nil {
		return Output{Error: err}
	}

	return Output{Result: user}
}

// FindByID func
func (u *UserRepositoryGorm) FindByID(id string) Output {
	var user models.User

	err := u.db.Where(&models.User{ID: id}).Take(&user).Error
	if err != nil {
		return Output{Error: err}
	}

	return Output{Result: &user}
}

// FindAll func
func (u *UserRepositoryGorm) FindAll() Output {
	var users models.Users

	err := u.db.Find(&users).Error
	if err != nil {
		return Output{Error: err}
	}

	return Output{Result: users}
}
