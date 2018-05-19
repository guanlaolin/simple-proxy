package main

import (
	"log"
	"net/http"
	"strings"
)

func http2httpsHandler(w http.ResponseWriter, r *http.Request) {
	//url := "https://" + r.Host + ":"
	//http.Redirect(w, r)
}

func domainHandler(w http.ResponseWriter, r *http.Request) {
	proxy := conf.Proxy

	host := strings.Split(r.Host, ":")[0]
	if "" == proxy[host] {
		log.Println("Proxy Not Found, Please check config")
		http.Error(w, "Proxy not set, Please contact admin.", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, proxy[host], http.StatusFound)
}
