package main

import (
	"io"
	"net"
	"time"
)

var conn net.Conn
var reader io.ReadCloser

func dial(netw, addr string) (net.Conn, error) {
	if conn != nil {
		conn.Close()
		conn = nil
	}
	netc, err := net.DialTimeout(netw, addr, 5*time.Second)
	if err != nil {
		return nil, err
	}

	conn = netc
	return netc, nil
}

func CloseCon() {
	if conn != nil {
		conn.Close()
	}
	if reader != nil {
		conn.Close()
	}
}
