package socket

import (
	"bufio"
	"net"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	// EventStatusBar reprensents the event for getting tmux statusbar.
	EventStatusBar = "status-bar\n"
)

// Pathname is the path where the socket is.
var Pathname = filepath.Join(os.TempDir(), "tline.sock")

// Request dials the socket with the given event.
func Request(event string) (string, error) {
	c, err := net.DialUnix("unix", nil, &net.UnixAddr{Name: Pathname, Net: "unix"})
	if err != nil {
		return "", errors.Wrap(err, "could not dial the socket")
	}
	_, err = c.Write([]byte(event))
	if err != nil {
		return "", errors.Wrap(err, "could not request the event")
	}

	reader := bufio.NewReader(c)
	return reader.ReadString('\n')
}

// Listen opens the socket and wait for incoming events.
func Listen(handler func(event string) string) error {
	ln, err := net.ListenUnix("unix", &net.UnixAddr{Name: Pathname, Net: "unix"})
	if err != nil {
		return errors.Wrap(err, "could not open socket")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		reader := bufio.NewReader(conn)
		if event, err := reader.ReadString('\n'); err == nil {
			payload := handler(event)
			_, _ = conn.Write([]byte(payload))
		}
	}

}
