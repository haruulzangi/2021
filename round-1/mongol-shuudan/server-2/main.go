package main

import (
	"bufio"
	"encoding/json"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const greeting = `Wassup!
Олегт шуудан хүргэж өгөхөд тус болооч?
Танд zipcode өгөхөд байршлын нэрийг хэлж өгч туслаарай.
`

func writeFully(conn net.Conn, data []byte) error {
	n := 0
	for n < len(data) {
		wrote, err := conn.Write(data[n:])
		if err != nil {
			return err
		}
		n += wrote
	}
	return nil
}

const (
	flag    = "HZ{97e5abe71_B@yr!a!aaa_P0STM2st3r_e7e7e7}"
	rounds  = 1000
	timeout = 5 * time.Second
)

func handleConnection(conn net.Conn, zipcodes map[string]string, codeList []string) {
	if err := writeFully(conn, []byte(greeting)); err != nil {
		conn.Close()
		return
	}

	scanner := bufio.NewScanner(conn)
	for i := 0; i < rounds; i++ {
		zipcode := codeList[rand.Intn(len(codeList))]
		if err := writeFully(conn, []byte(zipcode+" хаана байдаг вэ?\n")); err != nil {
			conn.Close()
			return
		}

		resultChan := make(chan bool)
		go func() {
			resultChan <- scanner.Scan()
		}()
		select {
		case ok := <-resultChan:
			if !ok {
				conn.Close()
				return
			}
			if strings.TrimSpace(scanner.Text()) != zipcodes[zipcode] {
				writeFully(conn, []byte("Буруу байна!\n"))
				conn.Close()
				return
			}
			if err := writeFully(conn, []byte("Nice!\n")); err != nil {
				conn.Close()
				return
			}
		case <-time.After(timeout):
			conn.Close()
			return
		}
	}

	writeFully(conn, []byte(flag+"\n"))
	conn.Close()
}

func main() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	zipcodes := make(map[string]string)
	if err := json.Unmarshal(data, &zipcodes); err != nil {
		panic(err)
	}

	codeList := make([]string, 0, len(zipcodes))
	for zipcode := range zipcodes {
		codeList = append(codeList, zipcode)
	}

	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn, zipcodes, codeList)
	}
}
