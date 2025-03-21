package main

import (
	"log"
	"os"
	"os/signal"
	"restfulapi/driver"
	"runtime/pprof"
	"syscall"
)

// Handle signal termination.
func handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	for sig := range c {
		log.Printf("Received signal %s", sig.String())
		signal.Stop(c)
		closeResources()
		os.Exit(15) // Exit as SIGTERM.
	}
}

// Global resources that must be closed
func closeResources() {
	childId := os.Getenv("CHILD_ID")
	if childId != "" {
		log.Printf("Shutting down children %s\n", childId)
	} else {
		log.Println("Shutting down parent")
	}
	if db := driver.ExportDbHandle(); db != nil {
		log.Println("Closing database connection")
		db.Close()
	}
	if cpuProfileFlag != "" && cpuProfileFile != nil {
		log.Println("Closing cpu profiling")
		pprof.StopCPUProfile()
		cpuProfileFile.Close()
	}
	if memProfileFlag != "" && memProfileFile != nil {
		log.Println("Closing memory profiling")
		memProfileFile.Close()
	}
	if childId != "" {
		log.Printf("Children %s shutdown\n", childId)
	} else {
		log.Println("Parent shutdown")
	}
}
