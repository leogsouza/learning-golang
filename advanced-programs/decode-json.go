// This sample program demonstrates how to decode a JSON string.
package main

import (
	"encoding/json" // Encoding and Decoding Package
	"fmt"
)

// JSON contains a sample String to unmarshal.

var JSON = `{
    "name": "Mark Taylor",
    "jobtitle": "Software Developer",
    "phone": {
        "home": "123-456-789",
        "office": "987-654-321"
    },
    "email": "john@doe.com"
}`

func main() {
	// Unmarshal the JSON string into info map variable
	var info map[string]interface{}
	if err := json.Unmarshal([]byte(JSON), &info); err != nil {
		panic(err)
	}

	// Print the output from info map.
	fmt.Println(info["name"])
	fmt.Println(info["jobtitle"])
	fmt.Println(info["email"])
	fmt.Println(info["phone"].(map[string]interface{})["home"])
	fmt.Println(info["phone"].(map[string]interface{})["office"])
}
