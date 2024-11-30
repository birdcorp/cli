package printer

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleAPIFailure(resp *http.Response) {
	if resp != nil {
		var errorResponse struct {
			Status string `json:"status"`
			Error  string `json:"error"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			log.Fatalf("Error decoding response: %v", err)
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case 404:
			log.Fatalf("%s: %s", errorResponse.Status, errorResponse.Error)
		case 400:
			log.Fatalf("%s: %s", errorResponse.Status, errorResponse.Error)
		case 401:
			log.Fatalf("%s: %s", errorResponse.Status, errorResponse.Error)
		case 403:
			log.Fatalf("%s: %s", errorResponse.Status, errorResponse.Error)
		default:
			log.Fatalf("%s: %s", errorResponse.Status, errorResponse.Error)
		}
	}
}
