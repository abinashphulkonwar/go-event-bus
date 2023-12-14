package client

import "bytes"

const scratchDefualtSize = 1024

// instace of client conected event bus
type Client struct {
	addr string

	buf bytes.Buffer
}

// create new client for event bus server
func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

// send message to event bus server
func (c *Client) Send(args []byte) error {
	_, err := c.buf.Write(args)
	if err != nil {
		return err
	}
	return nil
}

// recive message from event bus server
func (c *Client) Recive(scratch []byte) ([]byte, error) {

	if scratch == nil {
		scratch = make([]byte, scratchDefualtSize)
	}

	n, err := c.buf.Read(scratch)

	if err != nil {
		return nil, err
	}
	res := scratch[0:n]
	if len(res) == 0 {
		return res, nil
	}

	if res[len(res)-1] == '\n' {
		return res[:n-1], nil
	}
	
	lastPosition := bytes.LastIndexByte(res, '\n')
	println(lastPosition)
	return res[0:lastPosition], nil

}

func (c *Client) Connect() error {
	return nil
}

func (c *Client) Close() error {
	return nil
}
