package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	model "../Model"

	helper "../Helper"
	helperhttp "../Helper/Http"
)

//********** Data **********

const (
	ConfigFilename = "Configuration.json"
)

var defaultConfiguration = Data{
	API: httpServer{
		Host: "",
		Port: 8080,
		TimeoutSec: helperhttp.Timeout{
			Read:  15,
			Write: 15,
			Idle:  60,
		},
	},
	Monitoring: httpServer{
		Host: "",
		Port: 8081,
		TimeoutSec: helperhttp.Timeout{
			Read:  15,
			Write: 15,
			Idle:  60,
		},
	},
	DBConfig: model.DBConfig{
		Username: "root",
		Password: "rootroot",
		Host:     "127.0.0.1",
		Port:     3000,
		Name:     "elock",
		Driver:   "mysql",
	},
}

//********** Methods **********

type httpServer = helperhttp.ServerConfig

type Data struct {
	API        httpServer     `json:"api"`
	Monitoring httpServer     `json:"monitoring"`
	DBConfig   model.DBConfig `json:"dbConfig"`
}

func (obj *Data) read(filename string) (err error) {
	content, err := ioutil.ReadFile(filename)
	if err == nil {
		err = json.Unmarshal(content, obj)
	}
	return
}

func (obj Data) write(filename string) (err error) {
	content, err := json.Marshal(obj)
	if err == nil {
		err = ioutil.WriteFile(filename, content, 0644)
	}
	return
}

func ReadAndCreate(filename string) error {
	config := defaultConfiguration
	err := config.read(filename)

	fullpath, _ := filepath.Abs(filename)
	setConfig(fullpath, config)

	if os.IsNotExist(err) || err == nil {
		err = config.write(filename)
	}
	return err
}

func ReadSpringCloudConfig() error {
	springURI, found := os.LookupEnv("APP_SPRING_CONFIG_URI")
	if found == false {
		return fmt.Errorf("Spring cloud config uri not found")
	}
	profile, found := os.LookupEnv("APP_PROFILES")
	if found == false {
		return fmt.Errorf("Spring cloud config uri not found")
	}

	configPath := fmt.Sprintf("%v/%v-%v.json", springURI, Get().Name, profile)

	var restClient helper.RestClient
	content, statusCode, err := restClient.Request(configPath)
	if statusCode != 200 {
		return fmt.Errorf("Spring cloud config %v status code: %v", configPath, statusCode)
	} else if err != nil {
		return err
	}

	config := defaultConfiguration
	err = json.Unmarshal(content, &config)

	setConfig(configPath, config)

	return err
}
