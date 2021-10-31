package main

import (
	"bytes"
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

const flag = "HZ{ECDH_IS_EZ_4bb98f77d76170b1d8b3e7029299c3}"

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

func pcks7(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// ECDH on P384?
func handleConnection(conn net.Conn) {
	priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		panic(err)
	}

	if err := writeLPBytes(conn, priv.PublicKey.X.Bytes()); err != nil {
		return
	}
	if err := writeLPBytes(conn, priv.PublicKey.Y.Bytes()); err != nil {
		return
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

	shareX, _ := elliptic.P384().ScalarMult(peerX, peerY, priv.D.Bytes())
	sharedSecretBytes := sha256.Sum256(shareX.Bytes())
	sharedSecret := sharedSecretBytes[:]

	plaintext := pcks7([]byte(flag))
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		panic(err)
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		panic(err)
	}
	if err := writeLPBytes(conn, iv); err != nil {
		return
	}

	cbc := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(plaintext))
	cbc.CryptBlocks(ciphertext, plaintext)

	if err := writeLPBytes(conn, ciphertext); err != nil {
		return
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
