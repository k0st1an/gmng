package main

import (
	"io/ioutil"
	"log"

	y "gopkg.in/yaml.v2"
)

// Config ..
type Config struct {
	Region          string
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	Vault           string
}

var conf Config

func readConfig(file string) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	err = y.Unmarshal(bytes, &conf)
	if err != nil {
		log.Fatalln(err)
	}
}
