package server

import (
	"net"

	"github.com/xtlsoft/mascol/channel"
)

// Server is the server
type Server struct {
	Address string // Address to listen
	Context *Context
}

// Listen listens the address
func (s *Server) Listen() error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			// log err
		}
		go HandleConnection(conn, s.Context)
	}
}

// New makes a new server
func New(addr string) *Server {
	srv := new(Server)
	srv.Context = new(Context)
	srv.Context.Channels = make(map[string]*channel.Channel)
	srv.Address = addr
	return srv
}
