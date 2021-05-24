package helpers

import "os"

func GetEnv(key string, valueDefault string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return valueDefault
}
