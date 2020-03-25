package wzlib_crypto // Don't care about Go's "don't use underscores". Should be a better package names and read them nicer.

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	RSA_PEM_PUBKEY  = "public.pem"
	RSA_PEM_PRIVKEY = "private.pem"
	RSA_BIN_PUBKEY  = "public.key"
	RSA_BIN_PRIVKEY = "private.key"
)

type WzRSA struct {
	privKey *rsa.PrivateKey
	pubKey  *rsa.PublicKey
}

// NewWzRSA creates an instance of a class that takes care
// of keypair management operations (generation, rotation, encrytion etc).
func NewWzRSA() *WzRSA {
	wk := new(WzRSA)
	return wk
}

// GenerateKeyPair generates all public and private keys in PEM and Gob formats.
func (wk *WzRSA) GenerateKeyPair(pkiDir string) error {
	var err error
	wk.privKey, err = rsa.GenerateKey(rand.Reader, 0x800)
	if err != nil {
		return fmt.Errorf("Unable to generate keypair: %s", err.Error())
	}
	wk.pubKey = &wk.privKey.PublicKey

	if err := wk.saveGobKey(path.Join(pkiDir, RSA_BIN_PRIVKEY), wk.privKey); err != nil {
		return err
	}

	if err := wk.savePEMKey(path.Join(pkiDir, RSA_PEM_PRIVKEY), wk.privKey); err != nil {
		return err
	}

	if err := wk.saveGobKey(path.Join(pkiDir, RSA_BIN_PUBKEY), wk.pubKey); err != nil {
		return err
	}

	if err := wk.savePublicPEMKey(path.Join(pkiDir, RSA_PEM_PUBKEY), wk.pubKey); err != nil {
		return err
	}

	return nil
}

// LoadPEMKeyPair loads previously generated pub/priv keys
func (wk *WzRSA) LoadPEMKeyPair(pkiDir string) error {
	if err := wk.readPEMPrivateKey(path.Join(pkiDir, RSA_PEM_PRIVKEY)); err != nil {
		return err
	}
	if err := wk.readPEMPublicKey(path.Join(pkiDir, RSA_PEM_PUBKEY)); err != nil {
		return err
	}
	return nil
}

func (wk *WzRSA) saveGobKey(fileName string, key interface{}) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Unable to save Gob key: %s", err.Error())
	}
	defer outFile.Close()

	err = gob.NewEncoder(outFile).Encode(key)
	if err != nil {
		return fmt.Errorf("Unable to Gob-encode key: %s", err.Error())
	}
	return nil
}

func (wk *WzRSA) savePEMKey(fileName string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Unable to save PEM key: %s", err.Error())
	}
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	if err != nil {
		return fmt.Errorf("Unable to encode PEM key: %s", err.Error())
	}

	return nil
}

func (wk *WzRSA) savePublicPEMKey(fileName string, pubkey *rsa.PublicKey) error {
	asn1Bytes, err := asn1.Marshal(*pubkey)
	if err != nil {
		return fmt.Errorf("Unable to serialise public PEM key: %s", err.Error())
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Unable to create public PEM key file: %s", err.Error())
	}
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	if err != nil {
		return fmt.Errorf("Unable to encode public PEM key: %s", err.Error())
	}

	return nil
}

func (wk *WzRSA) readPEMPrivateKey(fileName string) error {
	priv, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes

	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return err
		}
	}
	wk.privKey, err = x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return err
	}

	return nil
}

func (wk *WzRSA) readPEMPublicKey(fileName string) error {
	pub, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes

	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return err
		}
	}

	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return err
	}

	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return err
	}
	wk.pubKey = key

	return nil
}
