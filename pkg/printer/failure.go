package printer

import (
	"log"
	"net/http"
)

func HandleAPIFailure(resp *http.Response) {
	if resp != nil {
		switch resp.StatusCode {
		case 404:
			log.Fatal("Resource not found")
		case 400:
			log.Fatal("Invalid request format")
		case 401:
			log.Fatal("Unauthorized: Please check your API credentials")
		case 403:
			log.Fatal("Forbidden: You don't have permission to access this resource")
		default:
			log.Fatalf("Error: HTTP %d", resp.StatusCode)
		}
	}
}
