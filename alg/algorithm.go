package alg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"log"
	"os"

	"JwsAlgsTest/model"
)

var (
	rsaKeys   model.RSA
	ecdsaKeys model.ECDSA
)

func GenRSAKeys() *model.RSA {
	err := error(nil)
	if rsaKeys.PrivateKey, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
		log.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}
	rsaKeys.PublicKey = &rsaKeys.PrivateKey.PublicKey

	return &rsaKeys
}

func GenECDSAKeys() *model.ECDSA {
	err := error(nil)
	if ecdsaKeys.PrivateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader); err != nil {
		log.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}

	ecdsaKeys.PublicKey = &ecdsaKeys.PrivateKey.PublicKey

	return &ecdsaKeys
}
