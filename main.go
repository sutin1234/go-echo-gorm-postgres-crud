package main

import (
	"fmt"
	"net/http"
	// _ "postgres_api/users/repository"
	routeuser "postgres_api/users/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)

func main() {
	fmt.Println("Golang GORM Started!")

	// var user models.User
	// db := database.InitailzeDB()
	// db.AutoMigrate(&user)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Echo Framework Go!")
	})

	uGroup := e.Group("/api/user")
	routeuser.RegisterRoute(uGroup)

	// userRepository := repository.NewUserRepositoryGorm(db)

	// e.GET("/api/user/", func(c echo.Context) error {
	// 	r := findUsers(userRepository)
	// 	return c.JSON(http.StatusOK, r)
	// })

	// e.GET("/api/user/add", func(c echo.Context) error {
	// 	insertUser(userRepository, "5")
	// 	return c.JSON(http.StatusOK, "Added!")
	// })

	// e.GET("/api/user/find/:id", func(c echo.Context) error {
	// 	r := findUser(userRepository, c.Param("id"))
	// 	return c.JSON(http.StatusOK, r)
	// })

	// e.GET("/api/user/delete/:id", func(c echo.Context) error {
	// 	id := c.Param("id")
	// 	deleteUser(userRepository, id)
	// 	return c.JSON(http.StatusOK, "Deleted!")
	// })

	// e.GET("/api/user/update/:id", func(c echo.Context) error {
	// 	updateUser(userRepository, c.Param("id"))
	// 	return c.JSON(http.StatusOK, "Updated!")
	// })

	// insertUser(userRepository, "3") // Insert
	// updateUser(userRepository) // Update
	// findUser(userRepository, "1") // FindByID
	// findUsers(userRepository) // FindUsers
	// deleteUser(userRepository, "2")

	// Start Services
	e.Logger.Fatal(e.Start(":8080"))
}

// func insertUser(r repository.UserRepository, id string) {
// 	u := &models.User{
// 		ID:       id,
// 		Name:     "thinny",
// 		LName:    "Injitt",
// 		Age:      30,
// 		Birthday: time.Now(),
// 		Email:    "tony.stin1234@gmail.com",
// 		Token:    "",
// 	}

// 	output := r.Save(u)
// 	if output.Error != nil {
// 		fmt.Println(output.Error.Error())
// 		os.Exit(1)
// 	}

// 	user, ok := output.Result.(*models.User)
// 	if !ok {
// 		fmt.Println("result is not a user")
// 	}

// 	fmt.Println(user)

// }

// func updateUser(r repository.UserRepository, id string) {
// 	output := r.FindByID(id)
// 	if output.Error != nil {
// 		fmt.Println(output.Error.Error())
// 		os.Exit(1)
// 	}

// 	user, ok := output.Result.(*models.User)
// 	if !ok {
// 		fmt.Println("result is not a user")
// 	}

// 	fmt.Println(user)
// 	// user.ID = "1"
// 	user.Name = "sutin modify"
// 	user.Age = 25

// 	fmt.Println(user)

// 	outputUpdate := r.Save(user)

// 	if outputUpdate.Error != nil {
// 		fmt.Println(outputUpdate.Error.Error())
// 		os.Exit(1)
// 	}

// 	userUpdated, ok := outputUpdate.Result.(*models.User)
// 	if !ok {
// 		fmt.Println("result is not a user")
// 	}

// 	fmt.Println(userUpdated)

// }

// func findUser(r repository.UserRepository, id string) *models.User {
// 	output := r.FindByID(id)
// 	if output.Error != nil {
// 		fmt.Println(output.Error.Error())
// 		os.Exit(1)
// 	}

// 	user, ok := output.Result.(*models.User)
// 	if !ok {
// 		fmt.Println("result is not a user")
// 	}

// 	fmt.Println(user)
// 	return user
// }

// func findUsers(r repository.UserRepository) models.Users {
// 	output := r.FindAll()
// 	if output.Error != nil {
// 		fmt.Println(output.Error.Error())
// 		os.Exit(1)
// 	}

// 	users, ok := output.Result.(models.Users)
// 	if !ok {
// 		fmt.Println("result is not a user")
// 	}

// 	fmt.Println(users)
// 	return users
// }

// func deleteUser(r repository.UserRepository, id string) {
// 	output := r.FindByID(id)
// 	if output.Error != nil {
// 		fmt.Println(output.Error.Error())
// 		os.Exit(1)
// 	}

// 	user, ok := output.Result.(*models.User)
// 	if !ok {
// 		fmt.Println("result is not a user")
// 	}

// 	user.ID = id

// 	outputDelete := r.Delete(user)
// 	if outputDelete.Error != nil {
// 		fmt.Println(output.Error.Error())
// 		os.Exit(1)
// 	}

// 	userDelete, ok := outputDelete.Result.(*models.User)
// 	if !ok {
// 		fmt.Println("result is not a user")
// 	}

// 	fmt.Println("user deleted")
// 	fmt.Println(userDelete)

// }
