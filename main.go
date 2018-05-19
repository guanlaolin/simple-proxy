// simple-proxy project main.go
package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"
)

const CONF_PATH = "./conf/config.json"

var waitGroup sync.WaitGroup
var conf *Config = new(Config)

func init() {
	// signal
	go signalProcess()

	// config
	err := conf.load(CONF_PATH)
	if err != nil {
		log.Fatalln("Config load failed:", err)
	}
	log.Println("Config load finished: ")
	conf.dispaly()
}

func main() {
	http.HandleFunc("/", domainHandler)

	waitGroup.Add(1)
	go listenHttpServer()

	waitGroup.Add(1)
	go listenHttpsServer()

	waitGroup.Wait()
	log.Println("Server exit...")
}

func listenHttpServer() {
	defer waitGroup.Done()

	log.Printf(" Addr:%s, Port:%d, Starting HTTP server...\n",
		conf.HTTP.Addr, conf.HTTP.Port)

	addr := conf.HTTP.Addr + ":" + strconv.Itoa(conf.HTTP.Port)
	log.Println("HTTP addr:", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("HTTP servr start failed,", err)
		return
	}

	log.Println("HTTP server exit...")
}

func listenHttpsServer() {
	waitGroup.Done()

	log.Printf(" Addr:%s, Port:%d, Starting HTTPS server...\n",
		conf.HTTPS.HTTPConfig.Addr, conf.HTTPS.HTTPConfig.Port)

	addr := conf.HTTPS.HTTPConfig.Addr + ":" +
		strconv.Itoa(conf.HTTPS.HTTPConfig.Port)
	log.Println("HTTPS addr:", addr)

	err := http.ListenAndServeTLS(addr,
		conf.HTTPS.Cert, conf.HTTPS.Key, nil)
	if err != nil {
		log.Println("Start HTTPS server failed:", err)
		return
	}

	log.Println("HTTP server exit...")
}
