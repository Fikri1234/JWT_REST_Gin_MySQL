package main

import (
	"JWT_REST_Gin_MySQL/configuration"
	"JWT_REST_Gin_MySQL/router"
	"JWT_REST_Gin_MySQL/util"
	"log"

	"github.com/spf13/viper"
)

func init() {

	// read config environment
	configuration.ReadConfig()

	util.Pool = util.SetupRedisJWT()

}

func main() {
	var err error

	// Setup database
	configuration.DB, err = configuration.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer configuration.DB.Close()

	// Setup router
	router := router.NewRoutes()

	port := viper.GetString("PORT")
	log.Fatal(router.Run(":" + port))
}
