package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func DigitalSignature () {
	//generate a new RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error in generating a private key", err)
		return
	}
	message := "Message to be signed"
	hashedMessage := sha256.Sum256([]byte(message))

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashedMessage[:])
	if err!= nil {
		fmt.Println("Error in signing the message: ", err);
		return
	}
	
	publicKey := &privateKey.PublicKey

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashedMessage[:], signature)
	if err!= nil {
		fmt.Println("Error in verifying the message: ", err)
		return
	}

	fmt.Println("Signature verified successfully")
}