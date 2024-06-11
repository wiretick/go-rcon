package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

const (
	// Packet types
	SERVERDATA_AUTH           int32 = 3
	SERVERDATA_AUTH_RESPONSE  int32 = 2
	SERVERDATA_EXECCOMMAND    int32 = 2
	SERVERDATA_RESPONSE_VALUE int32 = 0

	// Packet utils
	PKT_MIN_SIZE    = 10
	PKT_HEADER_SIZE = 8
)

type RCONServer struct {
	addr     string
	conn     net.Conn
	IsAuthed bool
}

func NewRCONServer(addr string) *RCONServer {
	return &RCONServer{
		addr: addr,
	}
}

func (s *RCONServer) Run() error {
	log.Println("Attempting to connect to server...")
	conn, err := net.Dial("tcp", "localhost:27015")
	if err != nil {
		return err
	}
	s.conn = conn

	log.Println("Successfully connected to server!")

	if err := s.Authenticate("123server"); err != nil {
		return err
	}

	return nil
}

func (s *RCONServer) Send(p *Packet) error {
	//data, _ := p.Encode()

	//fmt.Printf("Bytes to send: %v\n", data)

	if s.conn == nil {
		return fmt.Errorf("No connection")
	}
	s.conn.Write([]byte("some data to send"))
	return nil
}

func (s *RCONServer) Authenticate(password string) error {
	if err := s.Send(NewPacket(12, SERVERDATA_AUTH, password)); err != nil {
		return err
	}

	return nil
}

type Packet struct {
	Size int32 // Packet size field is excluded from size
	ID   int32
	Type int32
	Body string // null terminated x2 (x00)
}

func (p *Packet) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, p.Size))

	binary.Write(buf, binary.LittleEndian, p.Size)
	binary.Write(buf, binary.LittleEndian, p.ID)
	binary.Write(buf, binary.LittleEndian, p.Type)

	return []byte{}, nil
}

func (p *Packet) Decode(data []byte) error {
	buf := bytes.NewBuffer(data)

	binary.Read(buf, binary.LittleEndian, &p.Size)
	binary.Read(buf, binary.LittleEndian, &p.ID)
	binary.Read(buf, binary.LittleEndian, &p.Type)

	body := make([]byte, p.Size-PKT_HEADER_SIZE)
	buf.Read(body)
	p.Body = string(body)

	return nil
}

func NewPacket(id, ptype int32, body string) *Packet {
	size := PKT_MIN_SIZE + int32(len(body))
	body = body + "\x00\x00"

	return &Packet{
		Size: size,
		ID:   id,
		Type: ptype,
		Body: body,
	}
}

func main() {
	fmt.Println("GO RCON tool for Source protocol")

	newPacket := NewPacket(1, SERVERDATA_EXECCOMMAND, "sv_cheat 1")
	//	d, _ := newPacket.Encode()
	//fmt.Printf("data: %v", d)

	server := NewRCONServer("localhost:27015")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Packet: %+v\n", newPacket)
}
