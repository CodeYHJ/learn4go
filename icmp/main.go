package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ICMPType struct {
	Type     uint8
	Code     uint8
	CheckSum uint16
	Seq      uint16
}

func main() {
	fmt.Println("icmp")
}
func Struct2buf(data interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, &data)
	if err != nil {
		return nil, err
	}
	byt := buf.Bytes()
	buf.Reset()
	return byt, nil
}
func CheckSum(icmpByt []byte) uint16 {
	var (
		cSum   uint32 = 0
		index         = 0
		length        = len(icmpByt)
	)

	for length > 1 {
		cSum += uint32(icmpByt[index]) << 8
		cSum += uint32(icmpByt[index+1])
		index += 2
		length -= 2
	}
	//长度为基数
	if length > 0 {
		cSum += uint32(icmpByt[index])
	}
	cSum += cSum >> 16
	return uint16(^cSum)
}
