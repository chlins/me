package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/chlins/me/core"
)

func main() {
	reg := new(core.Request)
	reg.Role = "p"
	reg.Topic = "test"
	r, _ := json.Marshal(reg)
	conn, err := net.Dial("tcp", "127.0.0.1:6001")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r))
	conn.Write(r)
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		//fmt.Println(i)
		conn.Write([]byte("msg " + strconv.Itoa(i)))
	}
	fmt.Println("Cost time: ", time.Now().Sub(startTime))
}
