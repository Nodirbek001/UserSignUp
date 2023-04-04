package utils

import (
	"NewProUser/configs"
	"fmt"
)

func ConnectionURLBuilder(kind string) (string, error) {
	var config = configs.Load()
	var url string

	switch kind {
	case "postgres":
		url = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.PostgresHost,
			config.PostgresPort,
			config.PostgresUser,
			config.PostgresPassword,
			config.PostgresDatabase,
		)
	case "migration":
		url = fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			config.PostgresUser,
			config.PostgresPassword,
			config.PostgresHost,
			config.PostgresPort,
			config.PostgresDatabase,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", kind)
	}
	return url, nil
}
