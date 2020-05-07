package wzlib_crypto

import wzlib_logger "github.com/infra-whizz/wzlib/logger"

// WzClientCrypto class for all RSA/AES operations
type WzCryptoBundle struct {
	pair      *WzRSA
	symmetric *WzAES
	util      *WzCryptoUtils
	wzlib_logger.WzLogger
}

func NewWzCryptoBundle() *WzCryptoBundle {
	wcb := new(WzCryptoBundle)
	wcb.pair = NewWzRSA()
	wcb.symmetric = NewWzAES()
	wcb.util = NewWzCryptoUtils()

	return wcb
}

// InitPkiDir initialises the PKI directory with RSA keypair and AES token.
// Returns bool/bool, equivalent to rsa/aes accordingly.
func (wcb *WzCryptoBundle) InitPkiDir(pkiDir string) (bool, bool) {
	var err error
	var rsa bool = true
	var aes bool = true
	err = wcb.GetRSA().LoadPEMKeyPair(pkiDir)
	if err != nil {
		wcb.GetLogger().Errorf("Unable to load PKI directory: %s", err.Error())
		rsa = false
	}

	err = wcb.GetAES().LoadKey(pkiDir)
	if err != nil {
		wcb.GetLogger().Errorf("Unable to load AES token: %s", err.Error())
		aes = false
	}
	return rsa, aes
}

// GetRSA keypair API
func (wcb *WzCryptoBundle) GetRSA() *WzRSA {
	return wcb.pair
}

// GetAES token API
func (wcb *WzCryptoBundle) GetAES() *WzAES {
	return wcb.symmetric
}

// GetUtils returns crypto utils
func (wcb *WzCryptoBundle) GetUtils() *WzCryptoUtils {
	return wcb.util
}
