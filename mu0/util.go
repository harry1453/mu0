package mu0

func MachineCodeFromByteArray(data []byte) []uint16 {
	dataLength := len(data)
	if dataLength%2 != 0 {
		dataLength-- // Skip the last character
	}
	buffer := make([]uint16, dataLength/2)

	for i := 0; i < len(buffer); i++ {
		buffer[i] |= uint16(data[i]) & 0xFF
		buffer[i] <<= 8
		buffer[i] |= uint16(data[i]) & 0xFF
	}

	return buffer
}

func ByteArrayFromMachineCode(machineCode []uint16) []byte {
	buffer := make([]byte, len(machineCode)*2)

	for i := 0; i < len(machineCode); i++ {
		buffer[2*i] = uint8(machineCode[i] >> 8)
		buffer[2*i+1] = uint8(machineCode[i] & 0xFF)
	}

	return buffer
}
