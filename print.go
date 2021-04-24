package main

import (
	"encoding/json"
	"fmt"
)

func printJSON(i interface{}) error {
	bytes, err := json.Marshal(i)
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	return nil
}
