package main

func Uint16ToBytes(num uint16) []byte {
	return []byte{byte(num >> 8), byte(num & 0xFF)}
}
func Uint32ToBytes(num uint32) []byte {
	return []byte{byte(num >> 24), byte(num >> 16 & 0xFF), byte(num >> 8 & 0xFFFF), byte(num & 0xFFFFFF)}
}
