package env

import (
	"os"
	"strconv"
	"time"
)

// GetStr returns the value of an environment variable.
// If the environment variable is not set, it returns the default value.
func GetStr(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

// GetBool returns the value of an environment variable as a boolean.
// If the environment variable is not set or the value is not a valid boolean,
// it returns the default value (false).
func GetBool(key string, defaultValue bool) bool {
	v := GetStr(key, "")
	if v == "" {
		return defaultValue
	}

	b, err := strconv.ParseBool(v)
	if err != nil {
		return defaultValue
	}

	return b
}

// GetDuration returns the value of an environment variable as a time.Duration.
// If the environment variable is not set or the value is not a valid duration,
// it returns the default value (1 second).
func GetDuration(key string, defaultValue string) time.Duration {
	defv, err := time.ParseDuration(defaultValue)
	if err != nil {
		defv = time.Second
	}

	v := GetStr(key, "")
	if v == "" {
		return defv
	}

	d, err := time.ParseDuration(v)
	if err != nil {
		return defv
	}

	return d
}

// GetInt returns the value of an environment variable as an integer.
// If the environment variable is not set or the value is not a valid integer,
// it returns the default value (0).
func GetInt(key string, defaultValue int) int {
	v := GetStr(key, "")
	if v == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}

	return i
}
