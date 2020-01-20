package ehsan

import (
	"fmt"
	"net"
	"net/url"
	"time"
)

type NewServer struct {
	Addr        string
	MaxBodySize int64
	Routes      []Route
}

func (server *NewServer) Register(path string, handler Handler) {
	route := Route{
		Path:    path,
		Handler: handler,
	}
	server.Routes = append(server.Routes, route)
}

func (server *NewServer) Run() {

	// directly listen tcp
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("ehsan server started on : " + server.Addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			panic(err)
		}

		var bs = make([]byte, server.MaxBodySize)
		_, err = conn.Read(bs)
		if err != nil {
			conn.Write([]byte("Read error\n"))
		}

		if string(bs[0:3]) != "GET" {
			_, err = conn.Write([]byte("<h1>Method Not Allowed</h1>\n"))
			if err != nil {
				fmt.Println("internal server error")
			}
		}

		h1 := getURI(bs)

		query, err := url.Parse(h1)
		if err != nil {
			// os.Exit(1)
			// cmd.Exec("shutdown", "now")
			fmt.Println(err)
			conn.Write([]byte(err.Error() + "\n"))
		} else {
			fmt.Printf("[GET] Path : %v at %v \n", query.Path, time.Now().String())
			found := false
			for _, handler := range server.Routes {
				if handler.Path == query.Path {
					exec := handler.Handler(query)
					conn.Write(exec)
					found = true
				}
			}
			if !found {
				conn.Write([]byte("Not Found \n"))
			}
		}

		conn.Close()
	}
}
