package prettyprint

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/color"
)

// JSON takes an interface and prints it as a color-highlighted JSON string
func JSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	var prettyJSON interface{}
	err = json.Unmarshal(jsonData, &prettyJSON)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	keyColor := color.New(color.FgHiMagenta).SprintFunc()
	stringValueColor := color.New(color.FgGreen).SprintFunc()
	numberValueColor := color.New(color.FgMagenta).SprintFunc()
	boolValueColor := color.New(color.FgYellow).SprintFunc()
	nullValueColor := color.New(color.FgHiBlack).SprintFunc()

	printJSON(prettyJSON, keyColor, stringValueColor, numberValueColor, boolValueColor, nullValueColor, 0)
}

func printJSON(
	data interface{},
	keyColor, stringValueColor, numberValueColor, boolValueColor, nullValueColor func(a ...interface{}) string,
	indent int,
) {
	switch v := data.(type) {
	case map[string]interface{}:
		fmt.Println("{")
		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}
		for i, key := range keys {
			fmt.Printf("%s\"%s\": ", indentString(indent+2), keyColor(key))
			printJSON(v[key], keyColor, stringValueColor, numberValueColor, boolValueColor, nullValueColor, indent+2)
			if i < len(keys)-1 {
				fmt.Println(",")
			}
		}
		fmt.Printf("\n%s}", indentString(indent))
	case []interface{}:
		fmt.Println("[")
		for i, item := range v {
			fmt.Printf("%s", indentString(indent+2))
			printJSON(item, keyColor, stringValueColor, numberValueColor, boolValueColor, nullValueColor, indent+2)
			if i < len(v)-1 {
				fmt.Println(",")
			}
		}
		fmt.Printf("\n%s]", indentString(indent))
	case string:
		fmt.Printf("\"%s\"", stringValueColor(v))
	case float64:
		fmt.Printf("%s", numberValueColor(fmt.Sprintf("%.2f", v)))
	case bool:
		fmt.Printf("%s", boolValueColor(v))
	case nil:
		fmt.Printf("%s", nullValueColor("null"))
	default:
		fmt.Printf("\"%s\"", stringValueColor(v))
	}
}

func indentString(indent int) string {
	return fmt.Sprintf("%*s", indent, "")
}
