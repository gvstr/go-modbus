package modbus

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// FunctionCode 06 (0x06) Write Single Register
// This function code is used to write a single holding register in a remote device.
// The Request PDU specifies the address of the register to be written. Registers are addressed
// starting at zero. Therefore register numbered 1 is addressed as 0.
//
// The normal response is an echo of the request, returned after the register contents have
// been written.
//
// Quantity of Registers: 1
//
// Possible exception codes: 01, 02, 03 or 04
//
// Address range: 40000-49999
func (c *Client) WriteSingleRegister(startingAddress uint16, value int16) error {
	if err := validateInput(startingAddress, 1, 1); err != nil {
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
	if err := binary.Write(buffer, c.Byteorder, uint16(6)); err != nil {
		return err
	}
	// Unit identifier
	if err := binary.Write(buffer, c.Byteorder, c.UnitIdentifier); err != nil {
		return err
	}
	// Function code
	if err := binary.Write(buffer, c.Byteorder, uint8(6)); err != nil {
		return err
	}
	// Starting address
	if err := binary.Write(buffer, c.Byteorder, startingAddress); err != nil {
		return err
	}
	// Value
	if err := binary.Write(buffer, c.Byteorder, value); err != nil {
		return err
	}

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
