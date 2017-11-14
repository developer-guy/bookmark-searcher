package env

import "os"

func GetEnvironmentOrDefault(key string, fallback string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return fallback
	}
	return val
}
