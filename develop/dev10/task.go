package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func isCtrlD(b []byte) bool {
	return bytes.Equal([]byte{0x4}, b)
}

func main() {
	timeoutFlag := flag.String("timeout", "10s", "timeout value")

	flag.Parse()

	tf, _, _ := strings.Cut(*timeoutFlag, `s`)
	timeout, err := strconv.Atoi(tf)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//host, port := flag.Arg(0), flag.Arg(1)
	host, port := "localhost", "23"
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		signalName := <-sigChan
		fmt.Println("\nexiting with signal: ", signalName.String())
		conn.Close()
		os.Exit(0)

	}()

	//dataChan, start := make(chan string), make(chan bool)
	//go Reader(dataChan, start)
	//start <- true
	stdinReader := bufio.NewReader(os.Stdin)
	for {
		err = conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		buff := make([]byte, 1024)
		if _, err := conn.Read(buff); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(string(buff))

		input, err := stdinReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if isCtrlD(input) {
			conn.Close()
			fmt.Println("Conn closed!")
			os.Exit(0)
		}

		if _, err := conn.Write(input); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

}
