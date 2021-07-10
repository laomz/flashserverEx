package framework

import (
	"fmt"
	"net"
)

type TcpClient struct {
	atom
	conn net.Conn
	////////////////////////
	readBuff  []byte
	writeBuff chan []byte
}

func NewTcpClient() *TcpClient {
	return &TcpClient{
		readBuff:  make([]byte, MAX_PACK_LENGTH),
		writeBuff: make(chan []byte, MAX_PACK_LENGTH),
	}
}

func (tc *TcpClient) Read() error {
	return nil
}

func (tc *TcpClient) Write() error {
	return nil
}

func (tc *TcpClient) Close() error {
	tc.beActive.Set(false)
	if tc.conn != nil {
		tc.conn.Close()
	}
	return nil
}

func (tc *TcpClient) BeActive() bool {
	return tc.beActive.Get()
}

func (tc *TcpClient) GetAtomID() uint64 {
	return tc.id
}

func (tc *TcpClient) release() error {
	return nil
}

/********************************************************/
func (tc *TcpClient) StartTcpClient(id uint64, addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("TcpClient:StartTcpClient: err=%s", err.Error())
	}

	if err := tc.StartAtom(id, tc); err != nil {
		return fmt.Errorf("TcpClient:StartTcpPeer:err=%s", err.Error())
	}

	tc.conn = conn
	return nil
}

func (tc *TcpClient) Send(data []byte) error {
	return nil
}
