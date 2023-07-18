package getAddresses

import (
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
)

func GetAddress(attrName string) (address string) {
	outputs := make(map[string]map[string]interface{})
	content, err := os.ReadFile("../outputs.json")
	if err != nil {
		log.Println("Problem reading outputs.json")
		panic(err)
	}
	err = jsoniter.Unmarshal(content, &outputs)
	if err != nil {
		log.Println("Problem unmarshalling outputs.json")
		panic(err)
	}

	address = outputs[attrName]["value"].(string)

	return address
}
