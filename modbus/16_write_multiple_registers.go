package modbus

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// FunctionCode 16 (0x10) Write Multiple registers
// This function code is used to write a block of contiguous registers (1 to 123 registers) in a
// remote device.
//
// The requested written values are specified in the request data field. Data is packed as two
// bytes per register.
// The normal response returns the function code, starting address, and quantity of registers
// written
//
// Quantity of Registers: 1-123
//
// Possible exception codes: 01, 02, 03 or 04
//
// Address range: 40000-49999
func (c *Client) WriteMultipleRegisters(startingAddress uint16, values []uint16) error {
	if err := validateInput(startingAddress, uint16(len(values)), 123); err != nil {
		return err
	}

	// Buffer to hold the request message
	buffer := new(bytes.Buffer)

	// Transaction identifier
	if err := binary.Write(buffer, c.Byteorder, c.getTransactionIdentifier()); err != nil {
		return err
	}
	// Protocol identifier
	if err := binary.Write(buffer, c.Byteorder, uint16(0)); err != nil {
		return err
	}
	// Length
	if err := binary.Write(buffer, c.Byteorder, uint16(6+len(values))); err != nil {
		return err
	}
	// Unit identifier
	if err := binary.Write(buffer, c.Byteorder, c.UnitIdentifier); err != nil {
		return err
	}
	// Function code
	if err := binary.Write(buffer, c.Byteorder, uint8(16)); err != nil {
		return err
	}
	// Starting address
	if err := binary.Write(buffer, c.Byteorder, startingAddress); err != nil {
		return err
	}
	// Quantity of registers
	if err := binary.Write(buffer, c.Byteorder, uint16(len(values))); err != nil {
		return err
	}
	//Byte count
	if err := binary.Write(buffer, c.Byteorder, uint8(len(values))); err != nil {
		return err
	}
	// Value
	if err := binary.Write(buffer, c.Byteorder, values); err != nil {
		return err
	}

	b := buffer.Bytes()
	_ = b
	fmt.Printf("%x\n", buffer.Bytes())
	reply, err := c.sendRequest(buffer.Bytes())
	if err != nil {
		return err
	}

	if len(reply) <= 0 {
		return errors.New("device is in unknown state, did not receive a valid response after writing to register")
	}

	return nil
}
