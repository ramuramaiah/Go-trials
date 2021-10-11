package HashNCrypto

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func CRC_Test() {
	// create a hasher
	h := crc32.NewIEEE()
	// write our data to it
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Printf("crc is %d\n", v)
}

func getHash(filename string) (uint32, error) {
	// open the file
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// remember to always close opened files
	defer f.Close()
	// create a hasher
	h := crc32.NewIEEE()
	// copy the file into the hasher
	// - copy takes (dst, src) and returns (bytesWritten, error)
	_, err = io.Copy(h, f)
	// we don't care about how many bytes were written, but we do want to
	// handle the error
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}

func HashTest() {
	h1, err := getHash("C:/Work/test.txt")
	if err != nil {
		return
	}
	fmt.Printf("The hash is %d\n", h1)
}

func SHA1_Test() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Printf("The SHA1 value is %x\n", bs)
}
