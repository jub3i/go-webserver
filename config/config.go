package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func Load() error {
	godotenv.Load()

	required := [1]string{
		"GOWS_PORT",
	}

	for _, envvar := range required {
		_, exists := os.LookupEnv(envvar)
		if !exists {
			errmsg := "required environment variable `" + envvar + "` " +
				"unset, shutting down..."
			return errors.New(errmsg)
		}
	}

	return nil
}
