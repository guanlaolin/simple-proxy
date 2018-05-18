// simple-proxy project main.go
package main

import (
	"config"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const PATH_CONFIG = "/conf/config.json"

var waitGroup sync.WaitGroup

func init() {
	path := os.Getenv("SIMPLE_PROXY")
	log.Println("$SIMPLE_PROXY:", path)

	err := config.ConfigParse(path + PATH_CONFIG)
	if err != nil {
		log.Fatal("ConfigParse,", err)
	}
	fmt.Println("Read config done:", config.Config)
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

	httpConf, err := config.ConfigGetMapInterface("http")
	if err != nil {
		log.Println("Start HTTP server failed:", err)
		return
	}
	log.Println("HTTP config:", httpConf)

	addr, ok := httpConf["addr"].(string)
	if !ok {
		log.Println("HTTP server addr parse failed")
		return
	}

	port, ok := httpConf["port"].(float64)
	if !ok {
		log.Println("HTTP Server addr parse failed")
		return
	}

	addr += ":" + strconv.FormatFloat(port, 'f', -1, 64)
	log.Println("HTTP server addr:", addr)

	log.Println("Starting HTTP server...")
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("Start HTTP server failed,", err)
		return
	}
	log.Println("HTTP server exit...")
}

func listenHttpsServer() {
	waitGroup.Done()

	httpsConf, err := config.ConfigGetMapInterface("https")
	if err != nil {
		log.Println("Start HTTPS server failed:", err)
		return
	}
	log.Println("HTTPS config:", httpsConf)

	addr, ok := httpsConf["addr"].(string)
	if !ok {
		log.Println("HTTPS server addr parse failed")
		return
	}

	port, ok := httpsConf["port"].(float64)
	if !ok {
		log.Println("HTTPS Server addr parse failed")
		return
	}

	addr += ":" + strconv.FormatFloat(port, 'f', -1, 64)
	log.Println("HTTPS server addr:", addr)

	cert, ok := httpsConf["cert-path"].(string)
	if !ok {
		log.Println("HTTPS server cert path parse failed")
		return
	}

	key, ok := httpsConf["key-path"].(string)
	if !ok {
		log.Println("HTTPS server key path parse failed")
		return
	}

	log.Println("Starting HTTPS server...")
	err = http.ListenAndServeTLS(addr, cert, key, nil)
	if err != nil {
		log.Println("Start HTTPS server failed:", err)
		return
	}
	log.Println("HTTP server exit...")
}
