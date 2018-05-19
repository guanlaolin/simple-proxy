package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	conf.load(path)
}

func (conf *Config) dispaly() {
	fmt.Println(conf)
}
