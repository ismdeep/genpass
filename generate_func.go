package main

import (
	"crypto/rand"
	"github.com/ismdeep/args"
	"math/big"
)

// RandStr Generate Random String
func RandStr(base string, length int64) string {
	letters := []rune(base)
	lettersLen := new(big.Int).SetInt64(int64(len(letters)))
	results := make([]rune, length)
	for i := range results {
		v, _ := rand.Int(rand.Reader, lettersLen)
		results[i] = letters[v.Int64()]
	}
	return string(results)
}

// RandDigital Generate Random Digital
func RandDigital(length int64) string {
	return RandStr("0123456789", length)
}

// RandHex Generate Random Hex String
func RandHex(length int64) string {
	return RandStr("0123456789abcdef", length)
}

// RandDigitalAndAlphabet Generate Random String with digital and alphabet
func RandDigitalAndAlphabet(length int64) string {
	return RandStr("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", length)
}

// GenPass Generate Password
func GenPass(digital bool, lowerCase bool, upperCase bool, fuzzy bool, length int64) string {
	base := ""
	if digital {
		base += "0123456789"
	}
	if lowerCase {
		base += "abcdefghijklmnopqrstuvwxyz"
	}
	if upperCase {
		base += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if base == "" {
		base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*<>-_+=|"
	}

	if args.Exists("--without-fuzzy") {
		base = RemoveFuzzy(base)
	}

	return RandStr(base, length)
}
