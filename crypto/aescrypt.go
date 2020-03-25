package wzlib_crypto // Don't care about Go's "don't use underscores". Should be a better package names and read them nicer.

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path"
)

const AES_TOKEN = "token.aes"

type WzAES struct {
	key *[0x20]byte
}

func NewWzAES() *WzAES {
	wa := new(WzAES)
	return wa
}

// GenerateKey generates a random 256-bit key
func (wa *WzAES) GenerateKey(pkiDir string) error {
	wa.key = &[32]byte{}
	_, err := io.ReadFull(rand.Reader, wa.key[:])
	if err != nil {
		return err
	}
	buff := make([]byte, 0)
	for _, elm := range wa.key {
		buff = append(buff, elm)
	}
	if err := ioutil.WriteFile(path.Join(pkiDir, AES_TOKEN), buff, 0600); err != nil {
		return err
	}

	return nil
}

// LoadKey loads AES token from the pki directory.
func (wa *WzAES) LoadKey(pkiDir string) error {
	buff, err := ioutil.ReadFile(path.Join(pkiDir, AES_TOKEN))
	if err != nil {
		return err
	}
	if len(buff) != 0x20 {
		return fmt.Errorf("AES key length is not as expected: %d", len(buff))
	}
	wa.key = &[32]byte{}
	for idx, elm := range buff {
		wa.key[idx] = elm
	}

	return nil
}

// Encrypt encrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Output takes the
// form nonce|ciphertext|tag where '|' indicates concatenation.
func (wa *WzAES) Encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(wa.key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt decrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Expects input
// form nonce|ciphertext|tag where '|' indicates concatenation.
func (wa *WzAES) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(wa.key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil, ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():], nil)
}
