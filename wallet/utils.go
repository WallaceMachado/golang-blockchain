package wallet

import (
	"github.com/mr-tron/base58"
	"log"
)

func Base58Encode(input []byte) []byte {

	encode := base58.Encode(input) // feito poelo bytcoin com base 64  mesno 6 carasteres que podem ser confundidos

	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}

// 0 O l I + /
