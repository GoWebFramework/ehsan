package ehsan

import (
	"bytes"
	"io"
	"net"
)

type NewServer struct {
	Addr        string
	ReusePort   bool
	MaxBodySize int64
}

func (server *NewServer) Run() {
	// todo : validate ip:port format
	bindAddr := server.Addr

	if server.ReusePort {
		bindAddr += "?reuseport=true"
	}

	ln, err := net.Listen("tcp", bindAddr)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, server.MaxBodySize)

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		var buf bytes.Buffer
		io.Copy(&buf, conn)
		fmt.Println(buf)
	}
}
