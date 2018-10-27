package server

import "net"

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
	return nil
}
