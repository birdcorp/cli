package formatting

import (
	"log"
	"strconv"
)

func ParseFloat(value string) float64 {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatalf("Error parsing float: %v", err)
	}
	return result
}
