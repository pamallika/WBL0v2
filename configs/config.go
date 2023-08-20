package configs

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Nats_server struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Cluster_id string `yaml:"cluster_id"`
		Client_id  string `yaml:"client_id"`
		Channel    string `yaml:"channel"`
	} `yaml:"nats-server"`
	Database struct {
		Username   string `yaml:"user"`
		Password   string `yaml:"pass"`
		DBname     string `yaml:"dbname"`
		DriverName string `yaml:"driverName"`
	} `yaml:"database"`
	Http_server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http-server"`
}

func (cfg *Config) InitFile() {
	f, err := os.Open("../configs/config.yml")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
