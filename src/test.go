package main

import "./tide"

func dealClient(con *tide.Tconn) {
	defer con.Close()
	println("link a client")
	for {
		t_msg := con.Read()
		msg := string(t_msg)
		if msg == "ping1" || msg == "ping13" {
			con.Send([]byte("pong1+13"))
			println("send msg to client :" + msg)
		}
	}

}

func main() {
	server, err := tide.TListen("tcp", "127.0.0.1:8088")
	if err != nil {
		println(err)
		return
	}
	defer server.Close()
	println("the server run on 127.0.0.1:8088")
	for {
		con := server.Accept()
		if con == nil {
			continue
		}
		go dealClient(con)
	}
}
