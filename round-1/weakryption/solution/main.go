package main

import (
	"bufio"
	"encoding/hex"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(conn)
	for i := 0; i < 3; i++ {
		if !scanner.Scan() {
			panic(scanner.Err())
		}
	}

	dictionary := make(map[string]string)

	end := false
	block := []byte{0x00, 0x00}
	for !end {
		if block[0] == 0xFF && block[1] == 0xFF {
			end = true
		}

		if _, err := conn.Write([]byte("2\n")); err != nil {
			panic(err)
		}
		if !scanner.Scan() {
			panic(scanner.Err())
		}

		data := hex.EncodeToString(block)
		if _, err := conn.Write([]byte(data + "\n")); err != nil {
			panic(err)
		}

		if !scanner.Scan() {
			panic(scanner.Err())
		}
		encrypted := scanner.Text()
		if _, exists := dictionary[encrypted]; exists {
			log.Printf("Duplicate value for %s - %s", encrypted, data)
		} else {
			dictionary[encrypted] = data
		}

		for i := 0; i < 2; i++ {
			if !scanner.Scan() {
				panic(scanner.Err())
			}
		}

		if block[1] == 0xFF {
			block[0]++
			block[1] = 0x00
		} else {
			block[1]++
		}
	}

	if _, err := conn.Write([]byte("1\n")); err != nil {
		panic(err)
	}
	if !scanner.Scan() {
		panic(scanner.Err())
	}

	flagHex := ""
	encryptedFlag := scanner.Text()
	println(encryptedFlag)
	for i := 0; i < len(encryptedFlag); i += 4 {
		ciphertext := encryptedFlag[i : i+4]
		if value, ok := dictionary[ciphertext]; !ok {
			log.Panicf("Failed to find plaintext for %s", ciphertext)
		} else {
			log.Printf("%s -> %s", ciphertext, value)
			flagHex += value
		}
	}

	flag, _ := hex.DecodeString(flagHex)
	println(string(flag))
}
