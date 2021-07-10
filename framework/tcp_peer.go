package framework

import (
	"errors"
	"fmt"
	"net"
)

type TcpPeer struct {
	atom
	conn net.Conn
	////////////////////////
	readBuff  []byte
	writeBuff chan []byte
}

func NewTcpPeer() *TcpPeer {
	return &TcpPeer{
		readBuff:  make([]byte, MAX_PACK_LENGTH),
		writeBuff: make(chan []byte, MAX_PACK_LENGTH),
	}
}

func (tp *TcpPeer) Read() error {
	return nil
}

func (tp *TcpPeer) Write() error {
	return nil
}

func (tp *TcpPeer) Close() error {
	tp.beActive.Set(false)
	if tp.conn != nil {
		tp.conn.Close()
	}
	return nil
}

func (tp *TcpPeer) BeActive() bool {
	return tp.beActive.Get()
}

func (tp *TcpPeer) GetAtomID() uint64 {
	return tp.id
}

func (tp *TcpPeer) release() error {
	return nil
}

/********************************************************/
func (tp *TcpPeer) StartTcpPeer(id uint64, conn net.Conn) error {
	if conn == nil {
		return errors.New("TcpPeer:StartTcpPeer: conn = nil")
	}

	if err := tp.StartAtom(id, tp); err != nil {
		return fmt.Errorf("TcpPeer:StartTcpPeer:err=%s", err.Error())
	}

	tp.conn = conn
	return nil
}

/*
可能引发panic
*/
func (tp *TcpPeer) Send(data []byte) error {
	defer GuardFunction("TcpPeer:Send", tp.id)
	if !tp.BeActive() {
		return fmt.Errorf("TcpPeer:Send:closed,id=%+v", tp.id)
	}

	// t := time.NewTicker(time.Second)
	// t.Reset(0)
	// select {
	// case tp.writeBuff <- data:
	// case <-t.C:
	// 	return fmt.Errorf("TcpPeer:Send:timeout,id=%+v", tp.id)
	// }
	tp.writeBuff <- data
	return nil
}
