package modbus

import (
	"encoding/binary"
	"fmt"
	"net"
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
	UnitIdentifier        int
	transactionIdentifier uint16
}

func NewClient(name string, address string, port string, unitidentifier int, byteorder binary.ByteOrder) *Client {
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
	return c.Connection.Close()
}
