package p2p

import (
	"testing"

	_ "github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4333"
	tcpOpts := TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: NOPHandshakeFunc,
	}
	tr := NewTCPTransport(tcpOpts)
	assert.Equal(t, tr.ListenAddr, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
