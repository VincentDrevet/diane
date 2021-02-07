package configuration

import (
	"encoding/json"
	"io/ioutil"
)

//Configuration is struct that store setting of the application load from json file
type Configuration struct {
	DBPath      string `json:"dbpath"`
	DebugEnable bool   `json:"debug_enable"`
	DebugFile   string `json:"debug_file"`
}

//AppConfig is global variable that contain settings of the application
var AppConfig Configuration

//ReadConfigurationFile read json configuration file and return settings struct
func ReadConfigurationFile(jsonpath string) {
	content, err := ioutil.ReadFile(jsonpath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(content, &AppConfig)
}
