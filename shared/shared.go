package shared

import (
	"encoding/json"
	"log"
	"os"
)

func LoadFromFile(name string, structure any) error {
	data, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, structure)
}

func CheckError(err error) bool {
	if err != nil {
		log.Printf("error: %s", err)
		return true
	}
	return false
}

func CheckErrorP(err error) {
	if err != nil {
		log.Panicf("error: %s", err)
	}
}

func CheckErrorF(err error) {
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}