package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func http2httpsHandler(w http.ResponseWriter, r *http.Request) {
	https := r.URL

	https.Scheme = "https"
	https.Host = strings.Split(r.Host, ":")[0] + ":" + strconv.Itoa(conf.HTTPS.Port)
	log.Println("convert to https request:", https.String())

	http.Redirect(w, r, https.String(), http.StatusTemporaryRedirect)
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
