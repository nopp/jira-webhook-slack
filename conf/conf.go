package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	JiraDomain   string `json:"jiradomain"`
	SlackWebhook string `json:"slackwebhook"`
	IPPort       string `json:"ipport"`
}

// LoadConfiguration import conf from config.json
func LoadConfiguration() config {

	var config config

	appDirs, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	configFile, err := ioutil.ReadFile(appDirs + "/config.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(configFile, &config)
	return config
}
