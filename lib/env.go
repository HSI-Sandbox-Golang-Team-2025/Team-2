package lib

import (
	"os"
)

func GetEnv(key string, def string) string {
	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return def
}
