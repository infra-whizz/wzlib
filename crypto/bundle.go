package wzlib_crypto

import (
	"strings"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	wzlib_transport "github.com/infra-whizz/wzlib/transport"
)

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

// SignMessage signs all message content, return serialised byte array
func (wcb *WzCryptoBundle) SignMessage(msg *wzlib_transport.WzGenericMessage) ([]byte, error) {
	var buff strings.Builder
	buff.WriteString(msg.Payload[wzlib_transport.PAYLOAD_COMMAND].(string))
	buff.WriteString(msg.Jid)

	sig, err := wcb.GetRSA().Sign([]byte(buff.String()))
	if err != nil {
		return nil, err
	}

	msg.Payload[wzlib_transport.PAYLOAD_RSA_SIGNATURE] = sig

	return msg.Serialise()
}
