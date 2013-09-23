package main

import "tide"
import "time"

func dealClient(con *Tconn) {
	defer con.Close()
	println("link a client")
	for {
		t_msg := con.Read()
		msg = string(t_msg)
		println("read the client msg :" + msg)
		if msg == "ping1" || msg == "ping13" {
			con.Send([]byte("pong1+13"))
			println("send msg to client :" + msg)
		}
	}

}

func main() {
	server, err := TListen("127.0.0.1", 8088)
	if err != nil {
		println(err)
		return 0
	}
	defer server.Close()
	for {
		con := server.Accept()
		if con != nil {
			continue
		}
		go dealClient(con)
	}
}
