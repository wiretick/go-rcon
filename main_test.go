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
	t.Run("basic response test", func(t *testing.T) {
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

		d, _ := got.Encode()
		t.Errorf("%+v", d)
		t.Errorf("%+v", data)
	})

	t.Run("bit more special characters", func(t *testing.T) {
		data := []byte{
			0x4c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x55, 0x73, 0x61, 0x67,
			0x65, 0x3a, 0x20, 0x20, 0x6c, 0x6f, 0x67, 0x20,
			0x3c, 0x20, 0x6f, 0x6e, 0x20, 0x7c, 0x20, 0x6f,
			0x66, 0x66, 0x20, 0x3e, 0x0a, 0x63, 0x75, 0x72,
			0x72, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x20, 0x6c,
			0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x20, 0x74,
			0x6f, 0x3a, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x2c,
			0x20, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
			0x2c, 0x20, 0x75, 0x64, 0x70, 0x0a, 0x00, 0x00,
		}

		expected := NewPacket(
			0, SERVERDATA_RESPONSE_VALUE,
			"Usage:  log < on | off >\ncurrently logging to: file, console, udp\n",
		)
		got := &Packet{}

		if err := got.Decode(data); err != nil {
			t.Fatalf("Failed to decode: %v", err.Error())
		}

		if expected.Body != got.Body {
			t.Fatalf("Body is not the same: \nwant %q, \ngot  %q", expected.Body, got.Body)
		}

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("Failed to decode packet: got %+v, expected %+v", got, expected)
		}
	})

}
