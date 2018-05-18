package main

import (
	"config"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func http2httpsHandler(w http.ResponseWriter, r *http.Request) {
	//url := "https://" + r.Host + ":"
	//http.Redirect(w, r)
}

func domainHandler(w http.ResponseWriter, r *http.Request) {
	proxy, err := config.ConfigGetMapInterface("proxy")
	if err != nil {
		log.Println("ConfigGetMapInterface,", err)
	}

	host := strings.Split(r.Host, ":")[0]

	if nil == proxy[host] {
		log.Println("Proxy Not Found, Please check config")
		return
	}

	v, ok := proxy[host].(string)
	if !ok {
		log.Println(proxy[host], " is type ", fmt.Sprintf("%T", proxy[host]), " not type string")
	}

	http.Redirect(w, r, v, http.StatusFound)
}
