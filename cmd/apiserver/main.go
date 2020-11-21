package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/IFtech-A/http_rest_api/internal/app/apiserver"
	"gopkg.in/yaml.v2"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yaml", "path to config file")
}

func main() {
	flag.Parse()
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("Error while reading a file")
		return
	}

	config := apiserver.NewConfig()

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal("Error while parsing config file")
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatalf("Error on start %v", config.Store.DatabaseURL)
	}
}
