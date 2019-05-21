package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func ProtocolServer(port int) {
	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	var (
		buf = make([]byte, 8192)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)
LOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])
		switch err {
		case io.EOF:
			break LOOP
		case nil:
			dataSplited := strings.Fields(data)
			switch dataSplited[0] {
			case "getTokens":
				program := strings.Join(dataSplited[1:len(dataSplited)-1], " ")
				tokens := lexAnalyze(program)

				break
			}
			if isPacketEnd(data) {
				break LOOP
			}
		default:
			log.Fatalf("Receive data failed:%s", err)
			return
		}
	}
	w.Write([]byte("sai"))
	w.Flush()
}

func isPacketEnd(data string) (over bool) {
	over = strings.HasSuffix(data, "\rSTOPCONN\r")
	return
}
