package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nanobox-io/nanobox/util/proxy"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("please specify host and client connection credentials")
		fmt.Println("in the form of nanobox-proxy <host> <client>")
		fmt.Println("where <host> or client can look like 'tcp://:8080'")
		fmt.Println("example: nanobox-proxy tcp://:5678 unix:///tmp/socket.sock")
		return
	}

	hostNet, hostAddr := splitUri(os.Args[1])
	clientNet, clientAddr := splitUri(os.Args[2])
	fmt.Println(hostNet,hostAddr)
	fmt.Println(clientNet, clientAddr)
	err := proxy.Start(hostNet, hostAddr, clientNet, clientAddr)
	fmt.Println(err)
}


func splitUri(uri string) (string, string) {
	arr := strings.Split(uri, "://")
	if len(arr) != 2 {
		return "tcp", arr[0]
	}
	return arr[0], arr[1]
}