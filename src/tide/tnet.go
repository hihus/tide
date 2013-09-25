package tide

import (
	"errors"
	"io"
	"net"
)

const (
	MAXPACKSIAE = 1024
	HEADER      = 2
)

//net.Conn封装
type Tconn struct {
	maxPackSize int
	header      int
	con         net.Conn
	head        []byte
}

func (this *Tconn) Read() ([]byte, error) {
	if _, err := io.ReadFull(this.con, this.head); err != nil {
		return nil, err
	}
	size := getUint(this.head, this.header)
	buff := make([]byte, size)
	if size > this.maxPackSize {
		return nil, nil
	}
	if len(buff) == 0 {
		return nil, nil
	}
	_, err1 := io.ReadFull(this.con, buff)
	if err1 != nil {
		return nil, err1
	}
	return buff, nil
}

func (this *Tconn) ReadInto(buff []byte) []byte {
	if _, err := io.ReadFull(this.con, this.head); err != nil {
		return nil
	}
	var size = getUint(this.head, this.header)
	if size > this.maxPackSize {
		return nil
	}
	if len(buff) < size {
		buff = make([]byte, size)
	} else {
		buff = buff[0:size]
	}
	if len(buff) == 0 {
		return nil
	}
	// 不等待空消息
	if _, err := io.ReadFull(this.con, buff); err != nil {
		return nil
	}

	return buff
}
func Readline() {

}
func (this *Tconn) Send(msg []byte) error {
	size := len(msg)
	n_msg := make([]byte, size+this.header)
	setUint(n_msg, this.header, size)
	copy(n_msg[this.header:], msg)
	_, err := this.con.Write(n_msg)
	return err
}
func (this *Tconn) SendNoHeader(msg []byte) error {
	_, err := this.con.Write(msg)
	return err
}

func (this *Tconn) Close() {
	this.con.Close()
}

func NewTconn(conn net.Conn) *Tconn {
	return NewTconnWithOp(conn, MAXPACKSIAE, HEADER)
}

func NewTconnWithOp(conn net.Conn, pack, header int) *Tconn {
	return &Tconn{
		maxPackSize: pack,
		header:      header,
		con:         conn,
		head:        make([]byte, header),
	}
}

//TCP 封装
func TConnect(nettype, addr string) (*Tconn, error) {
	return TConnectWithOp(nettype, addr, HEADER, MAXPACKSIAE)
}

func TConnectWithOp(nettype, addr string, header, size int) (*Tconn, error) {
	conn, err := net.Dial(nettype, addr)
	if err != nil {
		return nil, errors.New("dail failed")
	}
	return NewTconnWithOp(conn, size, header), nil
}

type TListener struct {
	listener    net.Listener
	maxPackSize int
	header      int
}

func TListen(nettype, addr string) (*TListener, error) {
	return TListenWithOp(nettype, addr, HEADER, MAXPACKSIAE)
}

func TListenWithOp(nettype, addr string, header, size int) (*TListener, error) {
	listener, err := net.Listen(nettype, addr)
	if err != nil {
		return nil, errors.New("listen failed")
	}
	return &TListener{
		listener:    listener,
		maxPackSize: size,
		header:      header,
	}, nil
}

func (this *TListener) Accept() *Tconn {
	conn, err := this.listener.Accept()
	if err != nil {
		return nil
	}
	return NewTconnWithOp(conn, this.maxPackSize, this.header)
}

func (this *TListener) Close() error {
	return this.listener.Close()
}
