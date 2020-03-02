package module

import (
	"fmt"
	"postgres_api/users/models"
	"postgres_api/users/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

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
	u.Password = HashAndSaltPassword([]byte(u.Password))
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

// GenerateJWT func
func GenerateJWT(email, username string) (string, error) {
	mySecret := []byte(viper.GetString("app.secret"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["useremail"] = email + username
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 60 * 24).Unix

	tokenString, err := token.SignedString(mySecret)
	if err != nil {
		fmt.Errorf("token signed Error %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

// HashAndSaltPassword func
func HashAndSaltPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Sprintf("hash passwsord error %s", err)
	}
	return string(hash)
}

// ComparePassword func
func ComparePassword(hashPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

// GetLogin func
func GetLogin(r repository.UserRepository, body *models.User) (*models.User, error) {
	var err error
	// pwdHash := HashAndSaltPassword([]byte(body.Password))
	output := r.FindUserByUserName(body.UserName)
	if output.Error != nil {
		fmt.Println(output.Error.Error())
	}

	user, ok := output.Result.(*models.User)
	if !ok {
		fmt.Println("result is not a user")
	}

	if ok := ComparePassword(user.Password, []byte(body.Password)); ok {
		token, err := GenerateJWT(user.Email, user.Password)
		if err != nil {
			fmt.Println("result is not a user")
		}
		user.Token = token

		outputUpdate := r.Save(user)
		if outputUpdate.Error != nil {
			fmt.Println(output.Error.Error())
		}

		userUpdated, ok := outputUpdate.Result.(*models.User)
		if !ok {
			fmt.Println("result is not a user")
		}
		return userUpdated, nil
	}

	return nil, err

}
