package Config

import (
	"fmt"
	"os"
)

func DatabaseConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		os.Getenv("DB_CONNECTION"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))
}
