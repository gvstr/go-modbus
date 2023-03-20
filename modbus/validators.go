package modbus

import "fmt"

func validateInput(address uint16, numberOfRegisters uint16, maximumNumberOfRegisters uint16) error {
	if numberOfRegisters < 1 || numberOfRegisters > maximumNumberOfRegisters {
		return fmt.Errorf("number of registers has to be in the range of 1-%d", maximumNumberOfRegisters)
	}
	if address > 9999 {
		return fmt.Errorf("address (%d) cannot exceed 9999", address)
	}
	return nil
}
