package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	res, err := Parse("bellingcat.json")
	if err != nil {
		return
	}

	js, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	os.WriteFile("result.json", js, 0644)
}
