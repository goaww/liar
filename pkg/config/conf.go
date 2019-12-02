package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Conf struct {
	SendGrid struct {
		Key string `yaml:"key"`
	}
	Sender struct {
		Name string `yaml:"name"`
		Mail string `yaml:"mail"`
	}
}

func (c *Conf) LoadConf() *Conf {
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + "/configs/application.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Print("loadConf successful")
	return c
}
