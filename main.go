package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var flags struct {
	help   bool
	client bool
	server bool
	host   string
}

func init() {
	flag.BoolVar(&flags.help, "help", false, "Shows this help message")
	flag.BoolVar(&flags.client, "client", false, "Act as a TCP client")
	flag.BoolVar(&flags.server, "server", false, "Act as a TCP server")
	flag.StringVar(&flags.host, "host", "0.0.0.0:1234", "The host:port to connect/listen on")
}

func printConnInfoAndBlock(conn net.Conn) {
	b := make([]byte, 1)
	fmt.Printf("LocalAddr: %s\nRemoteAddr:%s\n\n", conn.LocalAddr(), conn.RemoteAddr())
	conn.Read(b)
}

func main() {
	flag.Parse()

	if flags.client {
		conn, err := net.Dial("tcp", flags.host)
		if err != nil {
			log.Fatal(err)
		}
		printConnInfoAndBlock(conn)
	} else if flags.server {
		l, err := net.Listen("tcp", flags.host)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(l.Addr())

		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go printConnInfoAndBlock(conn)
		}

	} else {
		flags.help = true
	}

	if flags.help {
		flag.PrintDefaults()
		return
	}

}
