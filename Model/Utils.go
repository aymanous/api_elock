package model

import (
	"encoding/json"
	"log"
)

func logObj(obj interface{}) {
	jsonData, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(string(jsonData))
	}
}

func logTxt(str interface{}) {
	log.Println(str)
}
