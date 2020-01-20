# ehsan

<img align="right" width="100" height="100" src="https://raw.githubusercontent.com/GoWebFramework/ehsan/master/ehsan.png">

(CURRENTLY) THIS IS JUST A TOY PROJECT 

## Installation

```bash
go get -u github.com/GoWebFramework/ehsan
```

## Usage

```go
package main

import "github.com/GoWebFramework/ehsan"

func hello(url ehsan.URI) []byte {
	return []byte("hello world")
}

func login(url ehsan.URI) []byte {
	return []byte("this is login page")
}

func main() {
	// default struct/value not available yet
	server := ehsan.NewServer{
		Addr: "127.0.0.1:3000",
		MaxBodySize: 50000, // in byte, will break request if req.len > MaxBodySize
	}

	// register your routes
	server.Register("/", hello)
	server.Register("/login", login)

	server.Run()
}
```