package modbus

// Function code 03 (0x03) Read Holding Registers
// This function code is used to read the contents of a contiguous block of holding registers in a
// remote device. The Request PDU specifies the starting register address and the number of
// registers. In the PDU Registers are addressed starting at zero. Therefore registers numbered
// 1-16 are addressed as 0-15.
//
// The register data in the response message are packed as two bytes per register, with the
// binary contents right justified within each byte. For each register, the first byte contains the
// high order bits and the second contains the low order bits.
//
// Quantity of Registers: 1-125
//
// Possible exception codes: 01, 02, 03 or 04
//
// Address range: 40000-49999
func (c *Client) ReadHoldingRegisters(startingAddress uint16, quantityOfRegisters uint16) ([]int16, error) {
	if err := validateInput(startingAddress, quantityOfRegisters, 125); err != nil {
		return nil, err
	}

	readRequest, err := c.getReadRequest(startingAddress, quantityOfRegisters, 3)
	if err != nil {
		return nil, err
	}

	reply, err := c.sendRequest(readRequest)
	if err != nil {
		return nil, err
	}

	result := bytesToInt16(reply[9:], c.Byteorder)

	return result, nil
}
