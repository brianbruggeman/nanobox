package proxy_test

import (
	"net"
	"io"
	"time"
	"testing"

	"github.com/nanobox-io/nanobox/util/proxy"
)

func TestProxyStart(t *testing.T) {
	go createEchoServer(t)

	go func() {
		err := proxy.Start("tcp", "127.0.0.1:6666", "tcp", "127.0.0.1:5555")
		t.Error("proxyStartErr: %s", err)
	}()

	<- time.After(time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		t.Error("unable to dial proxy")
	}

	toWrite := []byte("hello friend")
	written, err := conn.Write(toWrite)
	if err != nil || len(toWrite) != written {
		t.Error("failed to write")
	}

	toRead := make([]byte, written)
	read, err := conn.Read(toRead)
	if err != nil || read != written {
		t.Error("failed to read")
	}
	if string(toRead) != string(toWrite) {
		t.Error("information read not the same as what was written")
	}
}

func createEchoServer(t *testing.T)  {
	listener, err := net.Listen("tcp", "127.0.0.1:5555")
	if err !=  nil {
		t.Error("unable to start listener")
	}
	conn, err := listener.Accept()
	if err !=  nil {
		t.Error("unable to accept connections")
	}

	io.Copy(conn, conn)
	conn.Close()
	listener.Close()
}