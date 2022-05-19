package messages

import (
	"encoding/binary"
	"fmt"
	"net"

	"google.golang.org/protobuf/proto"
)

type MessageHandler struct {
	conn net.Conn
}

func NewMessageHandler(conn net.Conn) *MessageHandler {
	m := &MessageHandler{
		conn: conn,
	}

	return m
}

func (m *MessageHandler) RegistrationSuccess() {
	// m.conn.Write("success")
	fmt.Fprintf(m.conn, "success\n")
}

func (m *MessageHandler) RegistrationFail() {
	// m.conn.Write("fail")
	fmt.Fprintf(m.conn, "fail\n")
}

func (m *MessageHandler) Send(wrapper *Wrapper) error {
	serialized, err := proto.Marshal(wrapper)
	if err != nil {
		return err
	}

	prefix := make([]byte, 8)
	binary.LittleEndian.PutUint64(prefix, uint64(len(serialized)))
	m.conn.Write(prefix)
	m.conn.Write(serialized)
	return nil
}

func (m *MessageHandler) Sendout(aggdata *AggregateData) error {
	serialized, err := proto.Marshal(aggdata)

	if err != nil {
		return err
	}

	m.conn.Write(serialized)
	return nil
}

func (m *MessageHandler) Receive() (*Wrapper, error) {
	prefix := make([]byte, 8)
	m.conn.Read(prefix)

	payloadSize := binary.LittleEndian.Uint64(prefix)
	payload := make([]byte, payloadSize)

	total := uint64(0)
	for total < payloadSize {
		n, _ := m.conn.Read(payload[total:])
		total += uint64(n)
	}

	wrapper := &Wrapper{}
	err := proto.Unmarshal(payload, wrapper)
	return wrapper, err
}

func (m *MessageHandler) Close() {
	m.conn.Close()
}
