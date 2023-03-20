package modbus

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// Function code 05 (0x05) Write Single Coil
// This function code is used to write a single output to either ON or OFF in a remote device.
// The requested ON/OFF state is specified by a constant in the request data field. A value of
// FF 00 hex requests the output to be ON. A value of 00 00 requests it to be OFF. All other
// values are illegal and will not affect the output.
//
// The Request PDU specifies the address of the coil to be forced. Coils are addressed starting
// at zero. Therefore coil numbered 1 is addressed as 0. The requested ON/OFF state is
// specified by a constant in the Coil Value field. A value of 0XFF00 requests the coil to be ON.
// A value of 0X0000 requests the coil to be off. All other values are illegal and will not affect
// the coil.
//
// Quantity of Registers: 1
//
// Possible exception codes: 01, 02, 03 or 04
//
// Address range: 0000-9999
func (c *Client) WriteSingleCoil(startingAddress uint16, value bool) error {
	if err := validateInput(startingAddress, 1, 1); err != nil {
		return err
	}
	val := uint16(0x0000)
	if value {
		val = 0xFF00
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
	if err := binary.Write(buffer, c.Byteorder, uint8(5)); err != nil {
		return err
	}
	// Starting address
	if err := binary.Write(buffer, c.Byteorder, startingAddress); err != nil {
		return err
	}
	// Value
	if err := binary.Write(buffer, c.Byteorder, val); err != nil {
		return err
	}

	reply, err := c.sendRequest(buffer.Bytes())
	if err != nil {
		return err
	}

	if len(reply) <= 0 {
		return errors.New("device is in unknown state, did not receive a valid response after writing to coil")
	}

	return nil
}
