package modbus

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// FunctionCode 15 (0x0F) Write Multiple Coils
// This function code is used to force each coil in a sequence of coils to either ON or OFF in a
// remote device. The Request PDU specifies the coil references to be forced. Coils are
// addressed starting at zero. Therefore coil numbered 1 is addressed as 0.
//
// The requested ON/OFF states are specified by contents of the request data field. A logical '1'
// in a bit position of the field requests the corresponding output to be ON. A logical '0' requests
// it to be OFF.
//
// The normal response returns the function code, starting address, and quantity of coils forced.
//
// Quantity of Registers: 1968
//
// Possible exception codes: 01, 02, 03 or 04
//
// Address range: 0000-9999
func (c *Client) WriteMultipleCoils(startingAddress uint16, values []bool) error {
	if err := validateInput(startingAddress, uint16(len(values)), 1968); err != nil {
		return err
	}

	// Buffer to hold the request message
	buffer := new(bytes.Buffer)
	// Convert bools to bits
	valueBytes := boolsToBits(values)

	// Transaction identifier
	if err := binary.Write(buffer, c.Byteorder, c.getTransactionIdentifier()); err != nil {
		return err
	}
	// Protocol identifier
	if err := binary.Write(buffer, c.Byteorder, uint16(0)); err != nil {
		return err
	}
	// Length
	if err := binary.Write(buffer, c.Byteorder, uint16(6+len(valueBytes))); err != nil {
		return err
	}
	// Unit identifier
	if err := binary.Write(buffer, c.Byteorder, c.UnitIdentifier); err != nil {
		return err
	}
	// Function code
	if err := binary.Write(buffer, c.Byteorder, uint8(15)); err != nil {
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
	if err := binary.Write(buffer, c.Byteorder, uint8(len(valueBytes))); err != nil {
		return err
	}
	// Value
	if err := binary.Write(buffer, c.Byteorder, boolsToBits(values)); err != nil {
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
