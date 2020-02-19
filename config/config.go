package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

var required []string

func Init(r []string) (map[string]string, error) {
	required = r
	godotenv.Load()

	// the rest of this function makes sure required config is set
	for _, envName := range required {
		_, exists := os.LookupEnv(envName)
		if !exists {
			errmsg := "required environment variable `" + envName + "` " +
				"unset, shutting down..."
			return nil, errors.New(errmsg)
		}
	}

	return Get(), nil
}

func Get() map[string]string {
	env := make(map[string]string)
	for _, envName := range required {
		envVal, _ := os.LookupEnv(envName)
		env[envName] = envVal
	}

	return env
}
