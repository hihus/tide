package main

import "./tide"
import "time"

var str map[int]string = map[int]string{
	1:  "为撒开饭啦水电费凯撒和地方垃圾啊合适的话发神经的方式将对方拉萨的回复可拉伸的了客服哈萨克的话费卡拉是假的",
	2:  "大家好",
	3:  "我是hihu，你是谁？",
	4:  "大家真的很好吗",
	5:  "好骗吗，天下无贼啊",
	6:  "阿士大夫撒旦法师的法师的",
	7:  "啊上肯德基阿斯顿酸辣粉撒娇的功能",
	8:  "{阿道夫：sadf,sdf:asgs,sadgasd:sdf}",
	9:  "this is a beat",
	10: "clos sakdgasdfasdhfj jsadgfjasgdjfkajsskjdf",
	11: "hello world package",
	12: "跨世纪的发挥空间撒的发生",
	13: "按实际的发生",
	14: "123123124124124",
	15: "121企鹅2对得起我",
	16: "qq360腾讯电视卡京东方搜狗",
}

func send(con *tide.Tconn, out chan int) {
	for i := 1; i <= 15; i++ {
		con.Send([]byte(str[i]))
		println("send ths msg :" + str[i])
		time.Sleep(time.Second * 1)
	}
	out <- 1
}

func read(con *tide.Tconn) {
	for {
		t_msg, err := con.Read()
		if t_msg == nil && err != nil {
			println("disconnect the server !")
			return
		}
		msg := string(t_msg)
		println("client read the server msg:" + msg)
	}
}

func main() {
	var out chan int = make(chan int)
	cli, err := tide.TConnect("tcp", "127.0.0.1:8088")
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
