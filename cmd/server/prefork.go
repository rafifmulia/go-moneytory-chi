package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
)

// Customized from this https://github.com/TechEmpower/FrameworkBenchmarks/blob/master/frameworks/Go/chi/src/chi-sjson/prefork.go
// Why use prefork? Because chi performs better when use preforking mode. https://www.techempower.com/benchmarks/#hw=ph&test=query&section=data-r22&l=zijocf-cn3

func doPrefork(isChild bool, bind string) (listener net.Listener) {
	var err error
	var fl *os.File
	var tcplistener *net.TCPListener
	if !isChild {
		var addr *net.TCPAddr
		addr, err = net.ResolveTCPAddr("tcp", bind)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("TCP Address: %s\n", addr.String())
		tcplistener, err = net.ListenTCP("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		fl, err = tcplistener.File()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Shared TCP file:", fl.Name())
		children := make([]*exec.Cmd, runtime.NumCPU()/2)
		environ := os.Environ()
		environ = append(environ, "")
		for i := range children {
			n := len(os.Args)
			args := make([]string, n-1, n)
			copy(args, os.Args[1:])
			args = append(args, "-child")

			children[i] = exec.Command(os.Args[0], args...)
			children[i].Stdout = os.Stdout
			children[i].Stderr = os.Stderr
			children[i].ExtraFiles = []*os.File{fl}
			environ[len(environ)-1] = fmt.Sprintf("CHILD_ID=%d", i+1)
			children[i].Env = environ
			err = children[i].Start()
			if err != nil {
				log.Fatal(err)
			}
		}
		for _, ch := range children {
			err := ch.Wait()
			if err != nil {
				log.Printf("child error: %s", err)
			}
		}
		closeResources()
		os.Exit(0)
	} else {
		fl = os.NewFile(3, "")
		listener, err = net.FileListener(fl)
		if err != nil {
			log.Fatal(err)
		}
		runtime.GOMAXPROCS(runtime.NumCPU() / 2)
		log.Println("Run as a children", os.Getenv("CHILD_ID"))
	}
	return listener
}
