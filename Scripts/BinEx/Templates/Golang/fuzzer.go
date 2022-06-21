package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {

	var port, buffer int
	var postfix, prefix, ip string

	flag.StringVar(&ip, "ip", "192.168.180.131", "Specify target.")
	flag.IntVar(&port, "port", 31337, "Specify port.")
	flag.IntVar(&buffer, "buffer", 15000, "Specify buffer length. Default is 15000.")
	flag.StringVar(&prefix, "prefix", "", "Specify a prefix.")
	flag.StringVar(&postfix, "postfix", "\n", "Specify a postfix.")
	flag.Parse()

	//   s2 := strconv.Itoa(i)

	for i := 0; i < buffer; i++ {
		if i%50 != 0 {
			continue
		}
		conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
		fmt.Println("Fuzzing with", i, "bytes")
		if err != nil {
			log.Fatalf("[!] Error at offset %d: %s\n", i, err)
		}
		//bufio.NewReader(conn).ReadString('\n')

		data := "A"
		for n := 0; n <= i; n++ {
			data += "A"
		}

		fmt.Fprint(conn, prefix, data, postfix)
		bufio.NewReader(conn).ReadString('\n')
		//fmt.Fprint(conn, postfix)
		//bufio.NewReader(conn).ReadString('!')

		if err := conn.Close(); err != nil {
			log.Println("[!] Error at offset %d: %s\n", i, err)
		}
	}
}
