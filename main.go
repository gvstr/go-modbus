package main

import (
	"encoding/binary"
	"log"

	"github.com/gvstr/go-modbus/modbus"
)

func main() {
	plc := modbus.NewClient("TestPlc", "127.0.0.1", "502", 1, binary.BigEndian)
	err := plc.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer plc.Disconnect()
}
