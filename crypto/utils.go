package wzlib_crypto

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"log"
	"os"
	"strings"
)

type WzCryptoUtils struct{}

func NewWzCryptoUtils() *WzCryptoUtils {
	return new(WzCryptoUtils)
}

func (wcu *WzCryptoUtils) encodeDigest(digest *hash.Hash) string {
	var fp string
	for idx, ch := range hex.EncodeToString((*digest).Sum(nil)) {
		if idx%2 != 0 {
			fp += string(ch) + ":"
		} else {
			fp += string(ch)
		}
	}
	return strings.Trim(fp, ":")
}

// PEMKeyFingerprintFromBytes reads PEM key from an array of bytes and returns SHA256 fingerprint.
func (wcu *WzCryptoUtils) PEMKeyFingerprintFromBytes(key []byte) string {
	return wcu.PEMKeyFingerprintFromString(string(key))
}

// PEMKeyFingerprintFromString reads PEM key from a string and returns SHA256 fingerprint
func (wcu *WzCryptoUtils) PEMKeyFingerprintFromString(key string) string {
	digest := sha256.New()
	for _, line := range strings.Split(strings.TrimSpace(key), "\n") {
		cipherline := strings.TrimSpace(line) + "\n"
		if strings.Contains(cipherline, "PUBLIC KEY-----") {
			continue
		}
		_, err := digest.Write([]byte(cipherline))
		if err != nil {
			log.Printf("Error collecting SHA256 hash: %s", err.Error())
		}
	}
	return wcu.encodeDigest(&digest)
}

// PEMKeyFingerprintFromFile reads PEM key from a file and returns SHA256 fingerprint
func (wcu *WzCryptoUtils) PEMKeyFingerprintFromFile(keypath string) string {
	var fp string
	fh, err := os.Open(keypath)
	if err != nil {
		log.Printf("Unable to open PEM key file %s: %s\n", keypath, err.Error())
	} else {
		digest := sha256.New()
		defer fh.Close()
		scr := bufio.NewScanner(fh)
		for scr.Scan() {
			cipherline := scr.Text() + "\n"
			if strings.Contains(cipherline, "PUBLIC KEY-----") {
				continue
			}
			_, err := digest.Write([]byte(cipherline))
			if err != nil {
				log.Printf("Error collecting SHA256 hash: %s", err.Error())
			}
		}
		fp = wcu.encodeDigest(&digest)
	}

	return fp
}
