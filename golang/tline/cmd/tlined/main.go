package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mdouchement/tline/internal/pane"
	"github.com/mdouchement/tline/internal/socket"
	"github.com/sevlyar/go-daemon"
)

var (
	version  = "dev"
	revision = "none"
	date     = "unknown"
)

func main() {
	var dev bool
	flag.BoolVar(&dev, "dev", false, "dev mode - aka no daemon")
	var v bool
	flag.BoolVar(&v, "version", false, "print version")
	flag.Parse()

	if v {
		fmt.Printf("%s - build %.7s @ %s\n", version, revision, date)
		os.Exit(0)
	}

	if dev {
		worker()
	} else {
		daemonize()
	}
}

func daemonize() {
	cntxt := &daemon.Context{
		PidFileName: filepath.Join(os.TempDir(), "tline.pid"),
		PidFilePerm: 0644,
		LogFileName: "",
		LogFilePerm: 0640,
		WorkDir:     os.TempDir(),
		Umask:       027,
	}

	if len(daemon.ActiveFlags()) > 0 {
		d, err := cntxt.Search()
		if err != nil {
			log.Fatalln("Unable send signal to the daemon:", err)
		}
		daemon.SendCommands(d)
		return
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatalln(err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Println("- - - - - - - - - - - - - - -")
	log.Println("daemon started")

	go worker()

	err = daemon.ServeSignals()
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println("daemon terminated")
}

func worker() {
	os.Remove(socket.Pathname)

	socket.Listen(func(event string) string {
		switch event {
		case socket.EventStatusBar:
			return pane.StatusBar()
		}
		return "event not found"
	})
}
