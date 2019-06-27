package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s HASHSUM FILE\n", os.Args[0])
		os.Exit(1)
	}

	var hasher hash.Hash
	var method string

	switch len(os.Args[1]) {
	case 32:
		hasher = md5.New()
		method = "MD5"
	case 40:
		hasher = sha1.New()
		method = "SHA1"
	case 56:
		hasher = sha256.New224()
		method = "SHA224"
	case 64:
		hasher = sha256.New()
		method = "SHA256"
	case 96:
		hasher = sha512.New384()
		method = "SHA384"
	case 128:
		hasher = sha512.New()
		method = "SHA512"
	default:
		log.Fatalf("Could not recognize %s as hash", os.Args[1])
	}

	file, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatalf("Could not open file %s: %v", os.Args[2], err)
	}
	defer file.Close()

	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Fatalf("Could not hash file %s: %v", os.Args[2], err)
	}

	hashsum := hex.EncodeToString(hasher.Sum(nil))

	if hashsum == os.Args[1] {
		fmt.Printf("%s OK on file %s\n", method, os.Args[2])
	} else {
		fmt.Printf("%s FAILED on file %s\n", method, os.Args[2])
		os.Exit(2)
	}
}
