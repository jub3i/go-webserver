package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func Load() error {
	godotenv.Load()

	required := []string{
		"GOWS_PORT",
	}

	for _, element := range required {
		_, ok := os.LookupEnv(element)
		if !ok {
			errmsg := "required environment variable " +
				"`" + element + "` " +
				"unset, shutting down..."
			return errors.New(errmsg)
		}
	}

	return nil
}
