package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"net"
)

const (
	rounds     = 32
	key_size   = 32
	block_size = 2
)

func feistel(data, key []byte) []byte {
	if len(data) != block_size {
		panic("Invalid data size")
	}
	if len(key) != key_size {
		panic("Invalid key size")
	}

	var f = func(subblock, key byte) byte {
		return (subblock>>6 | key<<2) ^ (subblock<<2 | key>>6)
	}

	left, right := data[0], data[1]
	for i := 0; i < rounds; i++ {
		temp := right ^ f(left, key[i])
		right = left
		left = temp
	}

	return []byte{left, right}
}

func encryptData(plaintext, key []byte) []byte {
	if len(plaintext)%block_size != 0 {
		panic("Invalid data size")
	}
	if len(key) != key_size {
		panic("Invalid key size")
	}

	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += block_size {
		copy(ciphertext[i:], feistel(plaintext[i:i+block_size], key))
	}
	return ciphertext
}

const help = `1. Get encrypted flag
2. Encrypt some data
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

const flag = "HZ{USE_L0NG3R_BL0CKS_f5bd57fcd1874bab7b9cd800}"

func padding(plaintext []byte) []byte {
	if len(plaintext)%2 != 0 {
		return append(plaintext, 0x00)
	} else {
		return plaintext
	}
}

func handleConnection(conn net.Conn) {
	if err := writeFully(conn, []byte("Hi!\n")); err != nil {
		conn.Close()
		return
	}

	sessionKey := make([]byte, key_size)
	if n, err := rand.Read(sessionKey); n != key_size || err != nil {
		conn.Close()
		return
	}
	encryptedFlag := hex.EncodeToString(encryptData(padding([]byte(flag)), sessionKey))

	if err := writeFully(conn, []byte(help)); err != nil {
		conn.Close()
		return
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "1" {
			if err := writeFully(conn, []byte(encryptedFlag+"\n")); err != nil {
				break
			}
		} else if line == "2" {
			if err := writeFully(conn, []byte("Please enter your data:\n")); err != nil {
				break
			}
			if !scanner.Scan() {
				break
			}

			data, err := hex.DecodeString(scanner.Text())
			if err != nil {
				writeFully(conn, []byte("Invalid data, bye!\n"))
				break
			}

			ciphertext := encryptData(padding(data), sessionKey)
			if err := writeFully(conn, []byte(hex.EncodeToString(ciphertext)+"\n")); err != nil {
				break
			}
		} else {
			if err := writeFully(conn, []byte("What??\n")); err != nil {
				break
			}
		}

		if err := writeFully(conn, []byte(help)); err != nil {
			conn.Close()
			return
		}
	}
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}
