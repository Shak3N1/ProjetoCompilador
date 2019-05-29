package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
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
		buff = make([]byte, 8192)
		r    = bufio.NewReader(conn)
		w    = bufio.NewWriter(conn)
	)
LOOP:
	for {
		n, err := r.Read(buff)
		data := string(buff[:n])
		switch err {
		case io.EOF:
			break LOOP
		case nil:
			dataSplited := regexp.MustCompile(" ").Split(data, 2)
			switch dataSplited[0] {
			case "getTokens":
				tokens := lexAnalyze(dataSplited[1])
				for _, token := range tokens {
					w.Write([]byte(fmt.Sprintf("sendToken %s %s %d ", token.token,
						token.lexeme, token.line)))
				}
				w.Flush()
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
}

func isPacketEnd(data string) (over bool) {
	over = strings.HasSuffix(data, "\rSTOPCONN\r")
	return
}
