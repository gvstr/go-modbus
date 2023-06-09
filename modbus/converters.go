package modbus

import (
	"bytes"
	"encoding/binary"
)

func boolsToBits(values []bool) []byte {
	bytesRequired := uint(len(values)) / 8
	if len(values)%8 != 0 {
		bytesRequired++
	}
	result := make([]byte, bytesRequired)
	for i := 0; i < len(values); i++ {
		if values[i] { // if true at index
			result[i/8] |= (1 << (i % 8))
		}
	}
	return result
}

// i/8 is the index of the byte in the bytes slice that contains the bit for the current element.
// i%8 is the position of the bit within the byte that corresponds to the current element.
// 1<<(i%8) is a bit mask that has a 1 in the position corresponding to the current element.
// bytes[i/8]&(1<<(i%8)) is a bitwise AND operation that extracts the bit from the byte.
// If the result is not zero, then the bit is set, so result[i] is set to true.
func bytesToBools(bytes []byte) []bool {
	result := make([]bool, len(bytes)*8)
	for i := 0; i < len(result); i++ {
		if bytes[i/8]&(1<<(i%8)) != 0 {
			result[i] = true
		}
	}
	return result
}

func bytesToInt16(input []byte, byteOrder binary.ByteOrder) []int16 {
	result := make([]int16, len(input)/2)
	buf := bytes.NewReader(input)
	for i := 0; i < len(result); i++ {
		var n int16
		binary.Read(buf, byteOrder, &n)
		result[i] = int16(n)
	}
	return result
}
