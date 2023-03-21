package modbus

// Function code 02 (0x02) Read Discrete Inputs
// This function code is used to read from 1 to 2000 contiguous status of discrete inputs in a
// remote device. The Request PDU specifies the starting address, i.e. the address of the first
// input specified, and the number of inputs. In the PDU Discrete Inputs are addressed starting
// at zero. Therefore Discrete inputs numbered 1-16 are addressed as 0-15.
//
// The discrete inputs in the response message are packed as one input per bit of the data field.
// Status is indicated as 1= ON; 0= OFF. The LSB of the first data byte contains the input
// addressed in the query. The other inputs follow toward the high order end of this byte, and
// from low order to high order in subsequent bytes.
// If the returned input quantity is not a multiple of eight, the remaining bits in the final data byte
// will be padded with zeros (toward the high order end of the byte). The Byte Count field
// specifies the quantity of complete bytes of data.
//
// Quantity of Registers: 1-2000
//
// Possible exception codes: 01, 02, 03 or 04
//
// Address range: 10000-19999
func (c *Client) ReadDiscreteInputs(startingAddress uint16, quantityOfRegisters uint16) ([]bool, error) {
	if err := validateInput(startingAddress, quantityOfRegisters, 2000); err != nil {
		return nil, err
	}

	readRequest, err := c.getReadRequest(startingAddress, quantityOfRegisters, 2)
	if err != nil {
		return nil, err
	}

	reply, err := c.sendRequest(readRequest)
	if err != nil {
		return nil, err
	}

	result := bytesToBools(reply[9:])

	return result[:quantityOfRegisters], nil
}
