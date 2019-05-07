package base

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"strings"
	"time"
)

func ReadByte(reader *bytes.Reader) byte {
	b, _ := reader.ReadByte()

	return b
}

func ReadBytes(reader *bytes.Reader, length int) []byte {
	result := make([]byte, length)
	reader.Read(result)

	return result
}

func ReadWord(reader *bytes.Reader) uint16 {
	var uint16 word
	//word_byte := make([]byte, 2)
	reader.Read([]byte(word))

	return binary.BigEndian.Uint16([]byte(word))
}

func ReadDWord(reader *bytes.Reader) uint32 {
	dword_byte := make([]byte, 4)
	reader.Read(dword_byte)

	return binary.BigEndian.Uint32(dword_byte)
}

func ReadDWordL(reader *bytes.Reader) uint32 {
	dword_byte := make([]byte, 4)
	reader.Read(dword_byte)

	return binary.LittleEndian.Uint32(dword_byte)
}

func ReadQuaWord(reader *bytes.Reader) uint64 {
	qword_byte := make([]byte, 8)
	reader.Read(qword_byte)

	return binary.BigEndian.Uint64(qword_byte)
}

func ReadString(reader *bytes.Reader, length uint8) string {
	string_byte := make([]byte, length)
	reader.Read(string_byte)

	return string(string_byte)
}
