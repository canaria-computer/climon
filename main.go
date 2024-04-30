package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"slices"

	"github.com/farmergreg/rfsnotify"
)

func main() {
	watcher, err := rfsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	path := flag.String("path", "", "Watch directory")
	errOut := flag.Bool("err", true, "Whether to display an error")
	stdOut := flag.Bool("stdo", true, "Whether to display an standard output")
	flag.Parse()

	go func() {
		for {
			fmt.Println("================================================")
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Println("watcher.Events is not ok")
					return
				}
				log.Println("Update: ", event.Name)
				startSubp(*errOut, *stdOut)

			case err, ok := <-watcher.Errors:
				if !ok {
					log.Println("watcher.Errors is not ok")
					return
				}
				log.Println("error:", err)
			case <-quit:
				fmt.Println("^C Received")
				return
			}
		}
	}()

	if *path == "" {
		log.Fatal("path is empty")
	}
	err = watcher.Add(*path)
	if err != nil {
		log.Fatal(err)
	}

	<-quit
}

func startSubp(errOut bool, stdOut bool) {
	args := os.Args[1:]
	begin := slices.Index(args, "--")
	subpCmd := args[begin+1]
	subpArgs := args[begin+2:]

	out, err := exec.Command(subpCmd, subpArgs...).Output()
	if err != nil && errOut {
		log.Fatal("error:", err)
		os.Exit(1)
	}
	if stdOut {
		fmt.Printf("%s\n", out)
	}
}
