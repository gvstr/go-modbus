package modbus

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"
)

// No support for gateway functionality yet, need to split
// unitid out to a device type and have a slice of devices
// in client.
type Client struct {
	Name                  string
	Address               string
	Port                  string
	Byteorder             binary.ByteOrder
	Connection            net.Conn
	Pollrate              time.Duration
	UnitIdentifier        uint8
	transactionIdentifier uint16
	sync.Mutex
}

func NewClient(name string, address string, port string, unitidentifier uint8, byteorder binary.ByteOrder) *Client {
	return &Client{
		Name:                  name,
		Address:               address,
		Port:                  port,
		Byteorder:             byteorder,
		UnitIdentifier:        unitidentifier,
		transactionIdentifier: 0,
	}
}

func (c *Client) Connect() error {
	c.Lock()
	defer c.Unlock()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", c.Address, c.Port))
	if err != nil {
		return err
	}
	// conn.SetReadDeadline(time.Now().Add(socketReadTimeout))
	// conn.SetWriteDeadline(time.Now().Add(socketWriteTimeout))

	c.Connection = conn

	return nil
}

func (c *Client) Disconnect() error {
	c.Lock()
	defer c.Unlock()
	return c.Connection.Close()
}

func (c *Client) getTransactionIdentifier() uint16 {
	c.Lock()
	defer c.Unlock()
	if c.transactionIdentifier >= 65535 {
		c.transactionIdentifier = 0
	} else {
		c.transactionIdentifier += 1
	}
	return c.transactionIdentifier
}

func (c *Client) sendRequest(request []byte) ([]byte, error) {
	c.Lock()
	defer c.Unlock()
	if _, err := c.Connection.Write(request); err != nil {
		return nil, err
	}

	reply := make([]byte, 260)

	bytesRead, err := c.Connection.Read(reply)
	if err != nil {
		return nil, err
	}

	reply = reply[:bytesRead]

	return reply, nil
}
