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

	readHoldingRegisters, err := plc.ReadHoldingRegisters(0, 4)
	if err == nil {
		fmt.Println(readHoldingRegisters)
	}

	readInputRegisters, err := plc.ReadInputRegisters(0, 4)
	if err == nil {
		fmt.Println(readInputRegisters)
	}

	err = plc.WriteSingleCoil(0, true)
	if err != nil {
		fmt.Println(err)
	}

	err = plc.WriteSingleRegister(0, 123)
	if err != nil {
		fmt.Println(err)
	}

	err = plc.WriteMultipleCoils(10, []bool{true, true, true, true, true, true, true, true, false, true})
	if err != nil {
		fmt.Println(err)
	}
}
