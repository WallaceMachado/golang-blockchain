package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"log"
)

// chave privada deriva a chave publica e a chave publica deriva o address
const (
	checksumLength = 4
	version        = byte(0x00)
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey // algoritimo de cryptografia
	PublicKey  []byte
}

func (w Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)

	versionedHash := append([]byte{version}, pubHash...)
	checksum := Checksum(versionedHash)

	fullHash := append(versionedHash, checksum...)
	address := Base58Encode(fullHash)
	fmt.Printf("pub key: %x\n", w.PublicKey)
	fmt.Printf("pub hash: %x\n", pubHash)
	fmt.Printf("pub address: %x\n", address)
	return address
}

// gera a chave publica e privada
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256() // tipo da curva de 256bytes

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	// gera a chave publica
	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		log.Panic(err)
	}

	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

func Checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}
