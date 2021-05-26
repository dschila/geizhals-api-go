package helpers

import "os"

// To get the value from the environment with specific key.
// If no value is found with the provided key then the default value is returned.
func GetEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return defaultValue
}
