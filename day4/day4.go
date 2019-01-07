package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	key := "yzbqklnj"

	// calculate an md5 hash of a secret key (puzzle input) + a number
	// find the lowest positive number that produces the desired hash value
	// which is a value that starts with five zeros
	fiveZeros := getHashInputNumber(key, "00000")
	fmt.Println("Lowest number that has five zeros as a prefix:", fiveZeros)

	sixZeros := getHashInputNumber(key, "000000")
	fmt.Println("Lowest number that has six zeros as a prefix:", sixZeros)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func getHashInputNumber(key string, hashPrefix string) int {
	// calculate an md5 hash of a secret key (puzzle input) + a number
	// find the lowest positive number that produces the desired hash value
	// which is a value that starts with provided prefix
	inputNumber := 0
	for !(strings.HasPrefix(getMD5Hash(key+strconv.Itoa(inputNumber)), hashPrefix)) {
		inputNumber++
	}

	return inputNumber
}
