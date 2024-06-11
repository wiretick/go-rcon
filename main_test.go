package main

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	//packet := NewPacket(1, SERVERDATA_EXECCOMMAND, "sv_cheat 1")
	//toBe := []byte{20, 1, 2, 115, 118, 95, 99, 104, 101, 97, 116, 32, 49, 0, 0}

	//	data, err := packet.Encode()
	//if err != nil {
	//	t.Errorf("Failed to encode: %v", err.Error())
	//}

	//for i, v := range data {
	//	if v != toBe[i] {
	//		t.Errorf("Data is: %v, but expected: %v", v, toBe[i])
	//	}
	//}
}

func TestDecode(t *testing.T) {
	data := []byte{
		0x17, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x48, 0x4c, 0x53, 0x57,
		0x20, 0x3a, 0x20, 0x54, 0x65, 0x73, 0x74, 0x20,
		0x0a, 0x00, 0x00,
	}

	expected := NewPacket(0, SERVERDATA_RESPONSE_VALUE, "HLSW : Test \n")
	got := &Packet{}

	if err := got.Decode(data); err != nil {
		t.Errorf("Failed to decode: %v", err.Error())
	}

	if expected.Body != got.Body {
		t.Errorf("Failed to decode body: want %q, got %q", expected.Body, got.Body)
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Failed to decode packet: got %+v, expected %+v", got, expected)
	}
}
