package cp

import (
    "testing"
)

func TestHex2base64(t *testing.T) {
    d := []byte{ 0x4d, 0x61, 0x6e}
    b := Hex2base64(d)
    if b != "TWFu" {
        t.Error("Expected TWFu got ", b)
    }
}

func TestString2hex(t *testing.T) {
    d := "0f13"
    b, _ := String2hex(d)
    if len(b) != 2 || b[0] != 0x0f || b[1] != 0x13 {
        t.Error("Expected 0f13, for ", b)
    }
}

func TestEx1(t *testing.T) {
    data, _ := String2hex("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
    res := Hex2base64(data)
    if res != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
        t.Error("Expected SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t got ", res)
    }
}

func TestFixedxor(t *testing.T) {
    res, _ := Fixedxor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
    if res != "746865206b696420646f6e277420706c6179" {
        t.Error("Expected 746865206b696420646f6e277420706c6179 got ", res)
    }
}

func Testscore(t *testing.T) {
    res := score([]byte {97, 98, 99, 100, 95})
    if 4 != res {
        t.Error("Expected sore of 4, got ", res);
    }
}

func TestSinglebytexor(t *testing.T) {
    key, message := Singlebytexor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    if key != 88 {
       t.Error("Expected key 88, got ", key, " message ", message)
    }
}

func Testchallenge4(t *testing.T) {
}
