package module

import (
	"fmt"
	"postgres_api/users/models"
	"postgres_api/users/repository"

)

// FindUsers func
func FindUsers(r repository.UserRepository) models.Users {
	output := r.FindAll()
	if output.Error != nil {
		fmt.Println(output.Error.Error())
	}

	users, ok := output.Result.(models.Users)
	if !ok {
		fmt.Println("result is not a user")
	}
	return users
}

// FindUser func
func FindUser(r repository.UserRepository, id string) *models.User {
	output := r.FindByID(id)
	if output.Error != nil {
		fmt.Println(output.Error.Error())
	}

	user, ok := output.Result.(*models.User)
	if !ok {
		fmt.Println("result is not a user")
		return user
	}

	return user
}

// Add func
func Add(r repository.UserRepository, body *models.User) *models.User {

	u := &models.User{}
	u = body
	output := r.Save(u)
	if output.Error != nil {
		fmt.Println(output.Error.Error())
	}

	user, ok := output.Result.(*models.User)
	if !ok {
		fmt.Println("result is not a user")
	}
	return user

}

func UpdateUser(r repository.UserRepository, body *models.User) *models.User {
	output := r.FindByID(body.ID)
	if output.Error != nil {
		fmt.Println(output.Error.Error())
	}

	user, ok := output.Result.(*models.User)
	if !ok {
		fmt.Println("result is not a user")
	}

	user = body
	outputUpdate := r.Save(user)

	if outputUpdate.Error != nil {
		fmt.Println(outputUpdate.Error.Error())
	}

	userUpdated, ok := outputUpdate.Result.(*models.User)
	if !ok {
		fmt.Println("result is not a user")
	}

	return userUpdated

}

// DeleteUser func
func DeleteUser(r repository.UserRepository, id string) *models.User {
	output := r.FindByID(id)
	if output.Error != nil {
		fmt.Println(output.Error.Error())
	}

	user, ok := output.Result.(*models.User)
	if !ok {
		fmt.Println("result is not a user")
	}

	user.ID = id

	outputDelete := r.Delete(user)
	if outputDelete.Error != nil {
		fmt.Println(output.Error.Error())
	}

	userDelete, ok := outputDelete.Result.(*models.User)
	if !ok {
		fmt.Println("result is not a user")
	}

	return userDelete

}
