package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := Load()
	if err != nil {
		log.Fatal(err)
	}
}

var required = [3]string{
	"GOWS_PORT",
	"GOWS_SESSION_STORE_KEY",
	"GOWS_ENV",
}

func Load() error {
	godotenv.Load()

	// the rest of this function makes sure required config is set
	for _, envName := range required {
		_, exists := os.LookupEnv(envName)
		if !exists {
			errmsg := "required environment variable `" + envName + "` " +
				"unset, shutting down..."
			return errors.New(errmsg)
		}
	}

	return nil
}

func Get() (map[string]string, error) {
	env := make(map[string]string)
	for _, envName := range required {
		envVal, exists := os.LookupEnv(envName)
		if !exists {
			errmsg := "required environment variable `" + envName + "` " +
				"unset, shutting down..."
			return nil, errors.New(errmsg)
		}
		env[envName] = envVal
	}

	return env, nil
}
