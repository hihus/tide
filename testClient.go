package main

import "tide"
import "time"
import "strconv"

func send(con *Tconn) {
	for i := 0; i <= 20; i++ {
		con.Send([]byte("ping" + strconv.Itoa(i)))
		println("send ths msg :" + "ping" + strconv.Itoa(i))
		time.sleep(time.Second * 2)
	}
	out <- 1
}

func read(con *Tconn) {
	t_msg := con.Read()
	msg := string(t_msg)
	println("client read the server msg:" + msg)
}

func main() {
	var out chan int = make(chan int)
	cli, err := TConnect("127.0.0.1", 8088)
	if err != nil {
		println(err)
		return 0
	}
	defer cli.Close()
	go send(con, out)
	go read(con)
	<-out
}
