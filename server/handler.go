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
	var inChannels = []string{"_broadcast"}
	broadcastChannel := ctx.GetChannel("_broadcast")
	broadcastChannel.Connections = append(broadcastChannel.Connections, conn)
	var doClean = func() {
		for _, v := range inChannels {

		}
	}
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
			// Subscribe the channel
			if cmdMsg.Command == "subscribe" {
				name := cmdMsg.Parameters["name"]
				ch := ctx.GetChannel(name)
				ch.Connections = append(ch.Connections, conn)
			}
		}
		// Broadcast message
		if msg.Channel == "_broadcast" {

		}
	}
}
