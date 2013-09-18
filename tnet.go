package tide

import (
	"net"
	"errors"
	"io"
)

const(
	MAXPACKSIAE 1024
	HEADER 2
)
//net.Conn封装
type Tconn struct{
	maxPackSize int
	header int
	con net.Conn
	head []byte
}

func(this *Tconn) Read() []byte {
	if _, err := io.ReadFull(this.con,this.head); err != nil {
		return nil
	}
	size := getUint(this.head,this.header)
	if size > this.maxPackSize{
		return nil
	}
	buff := make([]byte,size)
	if len(buff) == 0{
		return nil
	}
	if _,err1 := io.ReadFull(this.con,buff);err1 != nil{
		return nil
	}
	return buff
}

func(this *Tconn) Send(msg []byte) error{
	_,err = this.con.Write(msg)
	return err
}

func(this *Tconn) Close(){
	this.con.Close()
}

func NewTconn(conn net.Conn){
	return newTconnWithOp(conn,MAXPACKSIAE,HEADER)
}

func NewTconnWithOp(conn net.Conn,pack,header int){
	return &Tconn{
		maxPackSize:pack,
		header:header,
		con:conn,
		head:make([]byte,header),
	}
}
//TCP 封装
func TConnect(nettype,addr string)(*Tconn,error){
	return TConnectWithOp(nettype,addr,HEADER,MAXPACKSIAE)
}

func TConnectWithOp(nettype,addr string,header,size int)(*Tconn,error){
	conn,err := net.Dial(nettype,addr)
	if err != nil{
		return nil,errors.New("dail failed")
	}
	return NewTconnWithOp(conn,size,header),nil
}

type TListener struct{
	listener net.Listener
	maxPackSize int
	header int
}

func TListen(nettype, addr string) (*TListener, error){
	return TListenWithOp(nettype,addr,HEADER,MAXPACKSIAE)
}

func TListenWithOp(nettype, addr string,header,size int) (*TListener, error){
	listener,err := net.Listen(nettype,addr)
	if err != nil{
		return nil,errors.New("listen failed")
	}
	return &TListener{
		listener:listener,
		maxPackSize:size,
		header:header,
	},nil
}

func(this *TListener)Accept() *Tconn {
	conn,err := this.listener.Accept()
	if err != nil{
		return nil
	}
	return NewTconnWithOp(conn,this.maxPackSize,this.header)
}

func(this *TListener)Close() error {
	return this.listener.Close()
}




