package modbus

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

func bytesToBools(bytes []byte) []bool {
	result := make([]bool, len(bytes)*8)
	for i := 0; i < len(result); i++ {
		if bytes[i/8]&(1<<(i%8)) != 0 {
			result[i] = true
		}
	}
	return result
}
