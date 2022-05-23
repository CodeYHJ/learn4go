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
	Seq      uint8
}

func main() {
	fmt.Println("icmp")
}
func struct2buf(data interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, &data)
	if err != nil {
		return nil, err
	}
	byt := buf.Bytes()
	return byt, nil
}
func checkSum(icmp ICMPType) {
	icmpByt, err := struct2buf(icmp)
	var checkSum uint32
	var index = 0
	var length = len(icmpByt)
	//长度为基数
	if length/2 > 0 {
		lastIcmpByt := icmpByt[length-1]
		sum := 0 + uint16(lastIcmpByt)
		checkSum += uint32(sum)
		icmpByt = icmpByt[:length-1]
		length = length - 1
	}
	for length > 0 {
		sum := uint16(icmpByt[index]) + uint16(icmpByt[index+1])
		checkSum += uint32(sum)
		index += 2
		length -= 2
	}

}
