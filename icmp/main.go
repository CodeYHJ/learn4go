package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type ICMPType struct {
	Type     uint8
	Code     uint8
	CheckSum uint16
	Seq      uint16
	Ident    uint16
}

func main() {
	ipaddr, err := net.ResolveIPAddr("ip", "baidu.com")
	if err != nil {
		fmt.Printf("ResolveIPAddr报错：%v \n", err)
	}
	fmt.Print(ipaddr.String(), "\n")
	for i := 1; i < 6; i++ {
		icmp := getICMP(uint16(i))
		if err = sendIcmp(icmp, ipaddr); err != nil {
			fmt.Print("fail sentICMP \n")
		}
		time.Sleep(2 * time.Second)
	}
}
func sendIcmp(icmp ICMPType, ipaddr *net.IPAddr) error {
	connect, err := net.DialIP("ip:icmp", nil, ipaddr)
	if err != nil {
		fmt.Printf("fail connect remote host: %s \n", err)
		return err
	}
	defer connect.Close()
	byt := WriteInBuf(icmp)
	_, err = connect.Write(byt)
	if err != nil {
		fmt.Printf("faile connect write icmp %v", err)
		return err
	}
	timeStart := time.Now()
	timeDead := time.Now().Add(time.Second * 2)
	if err := connect.SetReadDeadline(timeDead); err != nil {
		fmt.Printf("faile setReadDeadline: %v", err)
	}

	recv := make([]byte, 512)
	receive, err := connect.Read(recv)
	if err != nil {
		fmt.Printf("fail connect receive: %v", err)
		return err
	}
	timeEnd := time.Now()

	duration := timeEnd.Sub(timeStart).Nanoseconds() / 1e6

	fmt.Printf("%d bytes from %s: seq=%d ttl=%dms \n", receive, ipaddr.String(), icmp.Seq, duration)
	return nil
}
func WriteInBuf(icmp ICMPType) []byte {
	buffer := &bytes.Buffer{}
	err := binary.Write(buffer, binary.BigEndian, &icmp)
	if err != nil {
		fmt.Printf("failed binary.Write: %v", err)
		return buffer.Bytes()
	}
	defer buffer.Reset()
	return buffer.Bytes()
}
func getICMP(seq uint16) ICMPType {
	icmp := ICMPType{Type: 8, Code: 0, CheckSum: 0, Seq: seq, Ident: 0}
	byt := WriteInBuf(icmp)
	icmp.CheckSum = caculateCheckSum(byt)
	return icmp
}

func caculateCheckSum(icmpByte []byte) uint16 {
	var (
		checksum uint32 = 0

		length = len(icmpByte) - 1
	)
	for i := 0; i < length; i += 2 {
		sum := uint32(icmpByte[i])<<8 + uint32(icmpByte[i+1])
		checksum += sum
	}
	// 长度为基数
	if length&1 == 0 {
		checksum += uint32(icmpByte[length])
	}

	checksum = (checksum >> 16) + (checksum & 0xffff)
	checksum += checksum >> 16

	return uint16(^checksum)
}
