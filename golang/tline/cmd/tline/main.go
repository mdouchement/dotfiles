package main

import (
	"fmt"

	"github.com/mdouchement/tline/internal/socket"
)

var (
	version  = "dev"
	revision = "none"
	date     = "unknown"
)

func main() {
	payload, err := socket.Request(socket.EventStatusBar)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(payload)
	}
}
