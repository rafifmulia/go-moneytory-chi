package main

import (
	"log"
	"net"
	"net/http"
	"restfulapi/conf"
	"restfulapi/router"
)

var (
	_                      bool = conf.ParseFlag()
	preforkFlag, childFlag bool
	bindHost               string
)

func setFlags() {
	preforkFlag, childFlag = conf.GetPreforkFlag()
	bindHost = conf.GetBindHost()
	cpuProfileFlag = conf.GetCpuProfileFlag()
	memProfileFlag = conf.GetMemProfileFlag()
}

func initResources() {
	setFlags()
	profiling()
	conf.InitDbConnX()
}

func startListening(listener net.Listener) error {
	var err error
	if !preforkFlag {
		log.Printf("Listening on: %s\n", bindHost)
		err = http.ListenAndServe(bindHost, router.ChiInitRouter())
	} else {
		err = http.Serve(listener, router.ChiInitRouter())
	}
	return err
}

func main() {
	var listener net.Listener
	go handleSignals()
	initResources()
	if preforkFlag {
		listener = doPrefork(childFlag, bindHost)
	}
	if err := startListening(listener); err != nil {
		log.Fatal(err)
	}
}
