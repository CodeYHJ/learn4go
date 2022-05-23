package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestCheckSum(t *testing.T) {
	icmp := &ICMPType{Code: 0, Type: 8, CheckSum: 0, Seq: 1}
	byt, _ := Struct2buf(&icmp)
	checkSum := CheckSum(byt)
	icmp.CheckSum = checkSum

	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, icmp)
	if err != nil {
		panic(err)
	}
	byteAccpect := buf.Bytes()
	buf.Reset()
	var (
		length = len(byteAccpect)
		index  = 0
		sum    uint32
	)
	fmt.Println(length)
	for length > 1 {
		sum += uint32(byteAccpect[index]) << 8
		sum += uint32(byteAccpect[index+1])
		length -= 2
		index += 2
	}
	if length > 0 {
		sum += uint32(byteAccpect[index])
	}
	sum16 := uint16(^sum)
	if sum16 != 0 {
		t.Errorf("got %X, want %v", sum16, 0)
	}
}
