package main

import "./tide"

func dealClient(con *tide.Tconn) {
	defer con.Close()
	println("link a client")
	for {
		t_msg, err := con.Read()
		if t_msg == nil && err != nil {
			println("disconnect the client!")
			return
		}
		msg := string(t_msg)
		println(msg)
		if msg == "按实际的发生" {
			con.Send([]byte("pong!!!"))
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
