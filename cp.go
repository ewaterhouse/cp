package cp

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

func Hex2base64(h []byte) string {
	return base64.StdEncoding.EncodeToString(h)
}

func String2hex(s string) ([]byte, error) {
	h, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func Hex2string(h []byte) string {
	return hex.EncodeToString(h)
}

func Fixedxor(first, second string) (string, error) {
	if len(first) != len(second) {
		return "", errors.New("Unequal lengths")
	}
	h1, err := String2hex(first)
	if err != nil {
		return "", err
	}
	h2, err := String2hex(second)
	if err != nil {
		return "", err
	}
	res := make([]byte, len(first)/2)
	for k, _ := range h1 {
		res[k] = h1[k] ^ h2[k]
	}
	return Hex2string(res), nil
}

func score(candidate []byte) int {
	res := 0
	for _, v := range candidate {
		if (v >= 97 && (v <= 122)) || (v >= 65 && (v <= 90)) {
			res++
		}
		if v < 32 {
			res = res - 1
		}
	}
	return res
}

type try struct {
	key     byte
	message string
	score   int
}

func Singlebytexor(cypher string) (byte, int, string) {
	h, _ := String2hex(cypher)
	best := new(try)
	for i := byte(0); i < 128; i++ {
		attempt := make([]byte, len(h))
		for k, v := range h {
			attempt[k] = v ^ i
		}
		s := score(attempt)
		if s > best.score {
			best.key = i
			best.message = string(attempt)
			best.score = s
		}
		//fmt.Printf("key %v: score: %v %s\n", i, s, string(attempt))
	}
	return best.key, best.score, best.message
}

func Challenge4() error {
	file, err := os.Open("./4.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, score, message := Singlebytexor(scanner.Text())
		if score > 22 { //(len(scanner.Text())/2)-10 {
			fmt.Println(score, ":", scanner.Text(), " => ", message)
		}
	}
	return nil
}

func Repeatingxor(key string, plain string) (string, error) {
	hkey := []byte(key)
	hplain := []byte(plain)
	i := 0
	res := make([]byte, len(hplain))
	for k, v := range hplain {
		res[k] = v ^ hkey[i]
		i++
		if i == len(hkey) {
			i = 0
		}
	}
	return Hex2string(res), nil
}
