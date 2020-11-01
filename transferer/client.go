package transferer

import (
	"io/ioutil"
	"log"
	"net"
	"strconv"
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

func (c *Client) sendFileSize(path string) {
	size, err := GetSize(path)

	if err != nil {
		log.Fatalln(err)
	}

	c.conn.Write([]byte(strconv.Itoa(int(size))))
}

func (c *Client) sendFile(path string) {

	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%s size %d\n", path, len(content))
	c.conn.Write(content)
	log.Printf("%s send.\n", path)
}

func (c *Client) Run(path string) {

	defer c.conn.Close()
	c.sendFileName(path)
	c.sendFileSize(path)
	c.sendFile(path)

}
