package transferer

import (
	"log"
	"net"
)

type Server struct {
	listener net.Listener
	address  string
}

func NewServer(address string) (*Server, error) {
	server := Server{}
	err := server.init(address)
	return &server, err
}

func (s *Server) init(address string) error {

	var err error
	s.listener, err = net.Listen("tcp", address)
	s.address = address
	return err
}

func (s *Server) handleClient(conn net.Conn) {
	defer conn.Close()

	log.Println("client address ", conn.LocalAddr().String())
	conn.Write([]byte("This is a message\n"))
}

func (s *Server) Run() {
	defer s.listener.Close()
	log.Printf("Listening %s", s.address)

	for {
		conn, err := s.listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		go s.handleClient(conn)
	}

}
