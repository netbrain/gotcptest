package main

import (
	"net"
	"testing"
)

func TestTcpConnection(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:1234")
	go l.Accept()

	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Fatal(err)
	}
	if err := conn.Close(); err != nil {
		t.Fatal(err)
	}
}
