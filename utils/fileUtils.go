package utils

import (
	"encoding/json"
	"fmt"
	models "go-api/models"
	"io/ioutil"
	"log"
	"os"
)

func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}


func LoadKeys(db *models.Db) {
	keys, err := os.ReadFile("temp/data.json")
	if err != nil {
		log.Panic(err.Error())
	}
	json.Unmarshal(keys, &db)
	fmt.Println(keys)
}