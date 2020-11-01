package transferer

import (
	"log"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
}

func NewClient(serverAddress string) (*Client, error) {
	client := Client{}
	err := client.init(serverAddress)

	return &client, err
}

func (c *Client) init(address string) error {
	var err error
	c.conn, err = net.Dial("tcp", address)

	if err != nil {
		log.Println("Could not connect to the server")
		return err
	}

	log.Println("Connected to the server")
	return nil
}

func (c *Client) sendFileName(path string) {

	data := strings.Split(path, "/")
	c.conn.Write([]byte(data[len(data)-1]))
}

func (c *Client) Run(path string) {

	defer c.conn.Close()
	c.sendFileName(path)

}
