package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type HTTPConfig struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}

type HTTPSConfig struct {
	HTTPConfig
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

type Config struct {
	HTTP  HTTPConfig        `json:"http"`
	HTTPS HTTPSConfig       `json:"https"`
	Proxy map[string]string `json:"proxy"`
}

func (conf *Config) load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, conf)
}

func (conf *Config) reload(path string) {
	var temp *Config = new(Config)

	err := temp.load(path)
	if err != nil {
		log.Println("Config reload failed and without change. Please check config. ", err)
		return
	}

	log.Println("Config reload finished.")
	conf = temp
	conf.dispaly()
}

func (conf *Config) dispaly() {
	//fmt.Println(conf)
	fmt.Println("---------------- Current config ----------------")
	fmt.Println("HTTP: ", conf.HTTP)
	fmt.Println("HTTPS: ", conf.HTTPS)
	fmt.Println("Proxy: ")
	for k, v := range conf.Proxy {
		fmt.Printf("%s: %s\n", k, v)
	}
	fmt.Println("-----------------------END----------------------")
}
