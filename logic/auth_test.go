package logic

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_Rsa ...
func Test_Rsa(t *testing.T) {
	privKey, pubKey, err := GenRsaKey()
	assert.Nil(t, err)
	t.Logf("%s", string(privKey))
	t.Logf("%s", string(pubKey))

	data := []byte("This is a test.")
	sign, err := RsaSign(data, privKey)
	assert.Nil(t, err)
	t.Logf("sign: %s", hex.EncodeToString(sign))

	err = RsaVerify(data, sign, pubKey)
	assert.Nil(t, err)

	ciphertext, err := RsaEnc(data, pubKey)
	assert.Nil(t, err)
	t.Logf("ciphertext: %s", hex.EncodeToString(ciphertext))

	plaintext, err := RsaDec(ciphertext, privKey)
	assert.Nil(t, err)
	t.Logf("plaintext: %s", string(plaintext))
}
