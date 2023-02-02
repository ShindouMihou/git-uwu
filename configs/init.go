package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var RunningConfiguration = DefaultConfiguration

func Init() {
	bytes, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("configuration not found, using default configuration.")
		return
	}
	var configuration Config
	err = yaml.Unmarshal(bytes, &configuration)
	if err != nil {
		log.Fatal("failed to read configuration: ", err)
	}
	if configuration.Replacer == nil {
		configuration.Replacer = DefaultConfiguration.Replacer
	}
	if configuration.StutterChance == 0 {
		configuration.StutterChance = DefaultConfiguration.StutterChance
	}
	if configuration.Emojis == nil {
		configuration.Emojis = DefaultConfiguration.Emojis
	}
	RunningConfiguration = configuration
}
