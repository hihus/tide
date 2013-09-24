package main

import "./tide"
import "time"
import "strconv"

func send(con *tide.Tconn,out chan int) {
	for i := 0; i <= 20; i++ {
		con.Send([]byte("ping" + strconv.Itoa(i)))
		println("send ths msg :" + "ping" + strconv.Itoa(i))
		time.Sleep(time.Second * 2)
	}
	out <- 1
}

func read(con *tide.Tconn) {
	t_msg := con.Read()
	msg := string(t_msg)
	println("client read the server msg:" + msg)
}

func main() {
	var out chan int = make(chan int)
	cli, err := tide.TConnect("tcp","127.0.0.1:8088")
	if err != nil {
		println(err)
		return
	}
	defer cli.Close()
    println("this client runing ")
	go send(cli, out)
	go read(cli)
	<-out
}
