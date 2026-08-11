package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("../exercises/manifest.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	var exMap map[string][]string

	err = json.Unmarshal(file, &exMap)

	if err != nil {
		fmt.Println(err)
		return
	}

	for cat, topics := range exMap {
		fmt.Println(cat, topics)

		for _, topic := range topics {
			fmt.Println(topic)
		}

	}

}
