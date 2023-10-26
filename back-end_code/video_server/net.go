package main

import (
	"fmt"
	"net"
	"strings"
)

type MyNet struct {
	lis net.Listener
}

func (this *MyNet) netinit() {
	lis, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("net init errno")
	}
	this.lis = lis
	go this.lisrun()
}

func (this *MyNet) lisrun() {

	for {
		cnn, err := this.lis.Accept()
		if err != nil {
			fmt.Println("net init errno")
		}
		go this.conrun(cnn)
	}
}

func (this *MyNet) conrun(cnn net.Conn) {
	for {
		buf := make([]byte, 128)
		n, err := cnn.Read(buf)
		if n == 0 {
			fmt.Println("client drop")
		}
		if err != nil {
			fmt.Println("conn read errno")
			return
		}
		this.dosomething(string(buf[:n-1]))

	}
}

func (this *MyNet) dosomething(str string) {
	lei := strings.Split(str, "|")[0]
	fmt.Println(lei)
	//判断类型

}
