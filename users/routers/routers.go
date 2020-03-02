package routeuser

import (
	"fmt"
	"math/rand"
	"net/http"
	"postgres_api/database"
	"postgres_api/users/models"
	"postgres_api/users/module"
	"postgres_api/users/repository"
	"strconv"

	"github.com/labstack/echo/v4"

)

var user models.User
var db = database.InitailzeDB()

var userRepository = repository.NewUserRepositoryGorm(db)

// getInt func
func getInt(n int) int {
	return rand.Intn(n)
}

// RegisterRoute Group
func RegisterRoute(u *echo.Group) *echo.Group {

	db.AutoMigrate(&user)
	fmt.Println("tables users AutoMigrated")

	u.GET("", UserAll)
	u.GET("/", UserAll)
	u.GET("/:id", GetUser)
	u.POST("/add", AddUser)
	u.PUT("/:id", UpdateUser)
	u.DELETE("/:id", DeleteUser)
	return u
}

// UserAll func
func UserAll(c echo.Context) error {
	users := module.FindUsers(userRepository)
	return c.JSON(http.StatusOK, users)
}

// GetUser func
func GetUser(c echo.Context) error {
	id := c.Param("id")
	user := module.FindUser(userRepository, id)
	return c.JSON(http.StatusOK, user)
}

// AddUser func
func AddUser(c echo.Context) error {
	body := models.User{}
	body.ID = strconv.Itoa(getInt(100))
	if err := c.Bind(&body); err != nil {
		return err
	}
	user := module.Add(userRepository, &body)
	return c.JSON(http.StatusOK, user)
}

// UpdateUser func
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}

	body := models.User{}
	body.ID = id
	if err := c.Bind(&body); err != nil {
		return err
	}

	user := module.UpdateUser(userRepository, &body)
	return c.JSON(http.StatusOK, user)
}

// DeleteUser func
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}

	user := module.DeleteUser(userRepository, id)
	return c.JSON(http.StatusOK, user)
}
