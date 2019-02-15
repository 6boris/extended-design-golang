package util

import (
	"encoding/json"
	"fmt"
	"log"
)

func JsonMarshal(v interface{}) []byte {
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Println(err)
		log.Printf("json marshal errr: %s", err)
	}
	return data
}
