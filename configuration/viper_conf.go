package configuration

import (
	"os"

	"github.com/spf13/viper"
)

// ReadConfig default
func ReadConfig() {
	// default ==> setx APP_ENVIRONMENT STAGING
	if os.Getenv("APP_ENVIRONMENT") == "STAGING" {
		viper.SetConfigName("properties-staging")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./resource")
	} else if os.Getenv("APP_ENVIRONMENT") == "PROD" {
		viper.SetConfigName("properties-prod")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./resource")
	}

	viper.ReadInConfig()
}
