package config
import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load environment variables from.env file if it exists.
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading.env file")
		}
		fmt.Println("Loaded environment variables")

	} else {
		fmt.Println("app.env not found; skipping local environment loading.")
	}
}