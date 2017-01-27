package proxy

import (
	"net"
	"io"

	"github.com/jcelliott/lumber"
)


// establish a server and handle connections by piping them to the client
func Start(hostNet, hostAddr, clientNet, clientAddr string) error {
	listener, err := net.Listen(hostNet, hostAddr)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handleConnection(conn, clientNet,  clientAddr)
	}
}

func handleConnection(hostConn net.Conn, clientNet,  clientAddr string) {
	clientConn, err := net.Dial(clientNet, clientAddr)
	if err != nil {
		lumber.Error("%s", err)
		return
	}

	err = pipe(hostConn, clientConn)
	if err != nil {
		lumber.Error("%s", err)
		return
	}

	err = hostConn.Close()
	if err != nil {
		lumber.Error("%s", err)
		return
	}

	err = clientConn.Close()
	if err != nil {
		lumber.Error("%s", err)
		return
	}
}

func pipe(host, client net.Conn) error {
	go io.Copy(host, client)
	_, err := io.Copy(client, host)
	return err
}