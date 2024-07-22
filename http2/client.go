package http2

import (
	"fmt"
	"net"
)

var MaxFrameSize = 16384 // unless specified via SETTINGS_MAX_FRAME_SIZE

type HTTP2Client struct {
}

func (h *HTTP2Client) Request(method, host, payload string) (string, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return "", err
	}

	packet := []byte(fmt.Sprintf("%s %s %s\r\nHost: localhost:8080\r\n\r\n", method, host, Version))

	_, err = conn.Write(packet)
	if err != nil {
		return "", err
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func Connect(url string) (HTTP2Client, error) {
	return HTTP2Client{}, nil
}
