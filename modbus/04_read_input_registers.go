package modbus

import (
	"bytes"
	"encoding/binary"
)

// Function code 04 (0x04) Read Input Registers
// This function code is used to read from 1 to 125 contiguous input registers in a remote
// device. The Request PDU specifies the starting register address and the number of registers.
// In the PDU Registers are addressed starting at zero. Therefore input registers numbered 1-16
// are addressed as 0-15.
//
// The register data in the response message are packed as two bytes per register, with the
// binary contents right justified within each byte. For each register, the first byte contains the
// high order bits and the second contains the low order bits.
//
// Quantity of Registers: 1-125
// // Possible exception codes: 01, 02, 03 or 04
// // Address range: 30000-39999
func (c *Client) ReadInputRegisters(startingAddress uint16, quantityOfRegisters uint16) ([]int16, error) {
	if err := validateInput(startingAddress, quantityOfRegisters, 125); err != nil {
		return nil, err
	}

	// Buffer to hold the request message
	buffer := new(bytes.Buffer)

	// Transaction identifier
	if err := binary.Write(buffer, c.Byteorder, c.getTransactionIdentifier()); err != nil {
		return nil, err
	}
	// Protocol identifier
	if err := binary.Write(buffer, c.Byteorder, uint16(0)); err != nil {
		return nil, err
	}
	// Length
	if err := binary.Write(buffer, c.Byteorder, uint16(6)); err != nil {
		return nil, err
	}
	// Unit identifier
	if err := binary.Write(buffer, c.Byteorder, c.UnitIdentifier); err != nil {
		return nil, err
	}
	// Function code
	if err := binary.Write(buffer, c.Byteorder, uint8(4)); err != nil {
		return nil, err
	}
	// Starting address
	if err := binary.Write(buffer, c.Byteorder, startingAddress); err != nil {
		return nil, err
	}
	// Quantity of registers
	if err := binary.Write(buffer, c.Byteorder, quantityOfRegisters); err != nil {
		return nil, err
	}

	reply, err := c.sendRequest(buffer.Bytes())
	if err != nil {
		return nil, err
	}

	result := bytesToInt16(reply[9:], c.Byteorder)

	return result, nil
}
