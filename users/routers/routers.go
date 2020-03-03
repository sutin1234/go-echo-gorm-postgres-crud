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
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

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

	u.GET("", func(c echo.Context) error {
		c.Redirect(http.StatusSeeOther, "api/user/")
		return nil
	})
	u.GET("/", UserAll, AuthorizeMiddleware)
	u.GET("/:id", GetUser, AuthorizeMiddleware)
	u.POST("/add", AddUser)
	u.PUT("/:id", UpdateUser, AuthorizeMiddleware)
	u.DELETE("/:id", DeleteUser, AuthorizeMiddleware)
	u.POST("/login", LoginUser)

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
	_, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}

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
		return c.JSON(http.StatusNotFound, "user not found")
	}

	body := models.User{}
	body.ID = id
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusNotFound, "user not found")
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
		return c.JSON(http.StatusNotFound, "user not found")
	}

	user := module.DeleteUser(userRepository, id)
	return c.JSON(http.StatusOK, user)
}

// LoginUser func
func LoginUser(c echo.Context) error {
	body := models.User{}
	// body.ID = strconv.Itoa(getInt(100))
	if err := c.Bind(&body); err != nil {
		return err
	}

	userLogin, err := module.GetLogin(userRepository, &body)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	if userLogin == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	cookie := WriteCookie(user.Token)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, userLogin)
}

// WriteCookie func
func WriteCookie(value string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "jwt_token"
	cookie.Value = value
	cookie.Expires = time.Now().Add(24 * time.Hour)
	return cookie
}

// ReadCookie func
func ReadCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie("jwt_token")
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

// AuthorizeMiddleware func middleware
func AuthorizeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, map[string]string{
					// "error": fmt.Sprintf("panic: %s", r),
					"error": "Header Authorization Not Found",
				})
			}
		}()

		tokenString := c.Request().Header.Get("Authorization")[7:]

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Err sign method %v", token.Header["alg"])
			}
			return []byte(viper.GetString("app.secret")), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Token invalid")
		}

		return next(c)

	}
}
