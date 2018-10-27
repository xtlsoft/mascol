package channel

import "net"

// Channel is the channel structure
type Channel struct {
	Name        string           // Name of the channel
	Connections map[int]net.Conn // Connections of the channel
}
