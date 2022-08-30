package logic

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// GenRsaKey generate private key and public key
func GenRsaKey() ([]byte, []byte, error) {
	// generate private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}
	privStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privStream,
	}
	priv := pem.EncodeToMemory(block)

	// generate public key
	pubStream, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubStream,
	}
	pub := pem.EncodeToMemory(block)

	return priv, pub, nil
}

// RsaSign sign
func RsaSign(data, privKey []byte) ([]byte, error) {
	privateKey, err := decodePrivKey(privKey)
	if err != nil {
		return nil, err
	}

	digest := hashData(data)

	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest)
	if err != nil {
		return nil, err
	}
	return sign, nil
}

// RsaVerify verify sign
func RsaVerify(data, sign, pubKey []byte) error {
	publicKey, err := decodePubKey(pubKey)
	if err != nil {
		return err
	}
	digest := hashData(data)
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, digest, sign)
	if err != nil {
		return err
	}
	return nil
}

// RsaEnc encrypt
func RsaEnc(plaintext, pubKey []byte) ([]byte, error) {
	publicKey, err := decodePubKey(pubKey)
	if err != nil {
		return nil, err
	}
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// RsaDec decrypt
func RsaDec(ciphertext, privKey []byte) ([]byte, error) {
	privateKey, err := decodePrivKey(privKey)
	if err != nil {
		return nil, err
	}
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func decodePubKey(pubKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("invalid public key")
	}
	pubKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := pubKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not PublicKey type")
	}
	return publicKey, nil
}

func decodePrivKey(privKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privKey)
	if block == nil {
		return nil, errors.New("invalid private key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func hashData(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	digest := h.Sum(nil)
	return digest
}
