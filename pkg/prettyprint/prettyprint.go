package prettyprint

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/color"
)

func JSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	var prettyJSON map[string]interface{}
	err = json.Unmarshal(jsonData, &prettyJSON)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	keyColor := color.New(color.FgCyan).SprintFunc()
	valueColor := color.New(color.FgGreen).SprintFunc()

	printJSON(prettyJSON, keyColor, valueColor, 0)
}

func printJSON(data interface{}, keyColor func(a ...interface{}) string, valueColor func(a ...interface{}) string, indent int) {
	switch v := data.(type) {
	case map[string]interface{}:
		fmt.Println("{")
		for key, value := range v {
			fmt.Printf("%s\"%s\": ", indentString(indent+2), keyColor(key))
			printJSON(value, keyColor, valueColor, indent+2)
		}
		fmt.Printf("%s}", indentString(indent))
	case []interface{}:
		fmt.Println("[")
		for _, item := range v {
			printJSON(item, keyColor, valueColor, indent+2)
		}
		fmt.Printf("%s]", indentString(indent))
	case string:
		fmt.Printf("\"%s\",\n", valueColor(v))
	case float64:
		fmt.Printf("%s,\n", valueColor(v))
	case bool:
		fmt.Printf("%t,\n", v)
	case nil:
		fmt.Print("null,\n")
	default:
		fmt.Printf("\"%s\",\n", valueColor(v))
	}
}

func indentString(indent int) string {
	return fmt.Sprintf("%*s", indent, "")
}
