package authlib

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func ConvertStringToPrivateKey(privateKey string) (*rsa.PrivateKey) {
	block, err := pem.Decode([]byte(privateKey))
	if err != nil {
		panic(err)
	}
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}

func ConvertStringToPublicKey(publicKey string) (*rsa.PublicKey) {
	block, err := pem.Decode([]byte(publicKey))
	if err != nil {
		panic(err)
	}
	key, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	return key
}
