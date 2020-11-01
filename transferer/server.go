package transferer

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
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
	fileName, err := s.getFileName(conn)

	if err != nil {
		log.Println(err)
		return
	}

	size, _ := s.getSize(conn)
	fmt.Printf("%s = %d\n", fileName, size)
	s.receiveData(fileName, conn)

}

func (s *Server) receiveData(dstPath string, conn net.Conn) {
	log.Println("Receiving data :", dstPath)
	var buffer bytes.Buffer
	io.Copy(&buffer, conn)

	f, err := os.Create(dstPath)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Saving ..")
	f.Write(buffer.Bytes())
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

func (s *Server) getFileName(conn net.Conn) (string, error) {
	return s.readSmallBuffer(conn)
}

func (s *Server) getSize(conn net.Conn) (int, error) {

	buffer, err := s.readSmallBuffer(conn)

	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(buffer)

	return i, err
}

func (s *Server) readSmallBuffer(conn net.Conn) (string, error) {

	buffer := make([]byte, 256)
	n, err := conn.Read(buffer)

	if err != nil {
		return "", err
	}

	log.Println("")
	return string(buffer[:n]), nil
}
