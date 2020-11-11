package utils_BCCDP


import "bytes"

func PCKS5Padding(text []byte,blockSize int)[]byte{
	paddingSize :=blockSize - len(text)%blockSize
	paddingText :=bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
	return append(text, paddingText...)
}
func PCKS5RemovePadding(text []byte)[]byte  {
	paddingSize := text[len(text)-1]
	paddingText := bytes.Repeat([]byte{byte(paddingSize)},int(paddingSize))
	return bytes.TrimSuffix(text,paddingText)
}
