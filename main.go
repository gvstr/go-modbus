package main

import (
	"encoding/binary"
	"fmt"
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

	readCoils, err := plc.ReadCoils(0, 3)
	if err == nil {
		fmt.Println(readCoils)
	}

	readDiscreteInputs, err := plc.ReadDiscreteInputs(0, 4)
	if err == nil {
		fmt.Println(readDiscreteInputs)
	}
}
