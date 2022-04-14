package banana

import (
	"fmt"
	"os"
)

var endpoint string

func init() {
	// Set endpoint to whatever is in env
	switch env := os.Getenv("BANANA_URL"); env {
	case "":
		// Prod case, zero value
		endpoint = "http://api.banana.dev/"
	case "local":
		// Standard localdev case
		endpoint = "http://localhost/"
		fmt.Println("Running in devmode, endpoint base:", endpoint)
	default:
		// Custom localdev case
		endpoint = env
		fmt.Println("Running in devmode, endpoint base:", endpoint)
	}
}
