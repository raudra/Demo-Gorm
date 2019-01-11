package config


import (
	"os"
)

var GoEnv string

func init() {
	if os.Getenv("GO_ENV") == ""{
		os.Setenv("GO_ENV", "development")
	}
	GoEnv = os.Getenv("GO_ENV")
}

func IsGoenvDevelopment() bool{
	return GoEnv == "development"
}

func IsGoenvProduction() bool{
	return GoEnv == "production"	
}

func IsGoenvTest() bool{
	return GoEnv == "test"
}

