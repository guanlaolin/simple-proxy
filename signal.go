package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func signalProcess() {
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT)

	for {
		switch <-sig {
		case syscall.SIGHUP:
			log.Println("SIGHUP")
			conf.reload(CONF_PATH)
		case syscall.SIGINT:
			log.Println("SIGINT")
			resourcesClean()
		}
	}
}
