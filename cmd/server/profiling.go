package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	cpuProfileFlag string
	memProfileFlag string
	cpuProfileFile *os.File
	memProfileFile *os.File
)

// Profiling in this package is intended for
// enabled runtime profiling without http access such as
// import _ "net/http/pprof".
func profiling() {
	var err error
	if cpuProfileFlag != "" {
		cpuProfileFile, err = os.Create(cpuProfileFlag)
		if err != nil {
			log.Fatal(err)
		}
		runtime.SetCPUProfileRate(500)
		err = pprof.StartCPUProfile(cpuProfileFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	if memProfileFlag != "" {
		memProfileFile, err = os.Create(memProfileFlag)
		if err != nil {
			log.Fatal(err)
		}
		runtime.GC()
		err = pprof.WriteHeapProfile(memProfileFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}
