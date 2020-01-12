package envhelper

import (
	"fmt"
	"os"
)

// GetEnv simple helper function to read an environment or return a default value
func GetEnv(key string, defaultVal string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	return defaultVal, fmt.Errorf("The environment variable %s is not set; using default value of %s", key, defaultVal)
}

// GetRequiredEnv simple helper function to read an environment or die
func GetRequiredEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}

	return "", fmt.Errorf("The required environment variable %s is not set", key)
}
