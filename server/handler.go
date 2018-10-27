package server

import (
	"bufio"
	"io"
	"net"

	"github.com/xtlsoft/mascol/protocol"
)

// HandleConnection handles a connection
func HandleConnection(conn net.Conn, ctx *Context) {
	bio := bufio.NewReader(conn)
	for {
		line, err := bio.ReadBytes('\n')
		if err != nil {
			// When connection is closed
			if err == io.EOF {
				// TODO: Do some clean work...
			}
		}
		cmdMsg := new(protocol.CommandMessage)
		msg, err := protocol.ParseMessagePackage(line, cmdMsg)
		if err != nil {
			// TODO: need to add log
			conn.Close()
		}
		// Command message
		if msg.Channel == "_command" {

		}
		// Broadcast message
		if msg.Channel == "_broadcast" {

		}
	}
}
