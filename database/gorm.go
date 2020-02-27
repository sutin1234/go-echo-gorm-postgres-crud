package database

import (
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"

)

// GetGormConn func
func GetGormConn(host, user, dbname, password, port interface{}) (*gorm.DB, error) {
	return gorm.Open("postgres", fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s",
		host, port, user, dbname, password,
	))
}

// InitailzeDB func
func InitailzeDB() *gorm.DB {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("fatal error config file: %s \n", err))

	}

	host := viper.Get("db.host")
	user := viper.Get("db.user")
	password := viper.Get("db.password")
	port := viper.Get("db.port")
	dbname := viper.Get("db.dbname")

	db, err := GetGormConn(host, user, dbname, password, port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Database Connected!")
	return db

}
