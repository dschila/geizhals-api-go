package helpers

import "os"

// Environment variable value or default if not exists
func GetEnv(key string, valueDefault string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return valueDefault
}
