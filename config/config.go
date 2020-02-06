package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func Load() error {
	godotenv.Load()

	// the rest of this function makes sure config is set
	required := [2]string{
		"GOWS_PORT",
		"GOWS_SESSION_STORE_KEY",
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
