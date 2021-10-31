package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"math/big"
	"net"
)

func readLPBytes(conn net.Conn) ([]byte, error) {
	dataLenBuf := make([]byte, binary.MaxVarintLen32)

	n := 0
	for n < len(dataLenBuf) {
		read, err := conn.Read(dataLenBuf[n:])
		if err != nil {
			conn.Close()
			return nil, err
		}
		n += read
	}

	dataLen := int(binary.BigEndian.Uint32(dataLenBuf))

	data := make([]byte, dataLen)
	n = 0
	for n < dataLen {
		read, err := conn.Read(data[n:])
		if err != nil {
			return nil, err
		}
		n += read
	}

	return data, nil
}

func writeLPBytes(conn net.Conn, data []byte) error {
	dataLen := make([]byte, binary.MaxVarintLen32)
	binary.BigEndian.PutUint32(dataLen, uint32(len(data)))

	if _, err := conn.Write(dataLen); err != nil {
		conn.Close()
		return err
	}
	if _, err := conn.Write(data); err != nil {
		conn.Close()
		return err
	}

	return nil
}

func pcks7Decode(plaintext []byte) []byte {
	length := len(plaintext)
	padLength := int(plaintext[length-1])
	return plaintext[:(length - padLength)]
}

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		panic(err)
	}

	peerXBytes, err := readLPBytes(conn)
	if err != nil {
		return
	}
	peerYBytes, err := readLPBytes(conn)
	if err != nil {
		return
	}
	peerX := big.NewInt(0).SetBytes(peerXBytes)
	peerY := big.NewInt(0).SetBytes(peerYBytes)

	if err := writeLPBytes(conn, priv.X.Bytes()); err != nil {
		panic(err)
	}
	if err := writeLPBytes(conn, priv.Y.Bytes()); err != nil {
		panic(err)
	}

	shareX, _ := elliptic.P384().ScalarMult(peerX, peerY, priv.D.Bytes())
	sharedSecretBytes := sha256.Sum256(shareX.Bytes())
	sharedSecret := sharedSecretBytes[:]

	iv, err := readLPBytes(conn)
	if err != nil {
		panic(err)
	}
	ciphertext, err := readLPBytes(conn)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		panic(err)
	}

	plaintext := make([]byte, len(ciphertext))
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(plaintext, ciphertext)

	println(string(pcks7Decode(plaintext)))
}
