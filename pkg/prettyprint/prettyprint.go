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
		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}
		for i, key := range keys {
			fmt.Printf("%s\"%s\": ", indentString(indent+2), keyColor(key))
			printJSON(v[key], keyColor, valueColor, indent+2)
			if i < len(keys)-1 {
				fmt.Println(",")
			}
		}
		fmt.Printf("\n%s}", indentString(indent))
	case []interface{}:
		fmt.Println("[")
		for i, item := range v {
			printJSON(item, keyColor, valueColor, indent+2)
			if i < len(v)-1 {
				fmt.Println(",")
			}
		}
		fmt.Printf("\n%s]", indentString(indent))
	case string:
		fmt.Printf("\"%s\"", valueColor(v))
	case float64:
		fmt.Printf("%s", valueColor(fmt.Sprintf("%.2f", v)))
	case bool:
		fmt.Printf("%t", v)
	case nil:
		fmt.Print("null")
	default:
		fmt.Printf("\"%s\"", valueColor(v))
	}
}

func indentString(indent int) string {
	return fmt.Sprintf("%*s", indent, "")
}
