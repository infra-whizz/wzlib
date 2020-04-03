package wzlib_database_controller

import (
	wzlib_crypto "github.com/infra-whizz/wzlib/crypto"
	wzlib_transport "github.com/infra-whizz/wzlib/transport"
)

// WzClient entity object for the database
type WzClient struct {
	ID     int    `gorm:"primary_key"`
	Uid    string `gorm:"unique; not null"`
	Fqdn   string `gorm:"unique; not null"`
	RsaPk  string `gorm:"unique; not null"`
	RsaFp  string
	Status int `gorm:"not null"`
}

// NewWzClient creates an instance of the WzClient
func NewWzClient() *WzClient {
	wcc := new(WzClient)
	return wcc
}

// NewWzClientFromPayload creates an instance of the WzClient and fills-in with the payload
func NewWzClientFromPayload(payload map[string]interface{}) *WzClient {
	wcc := new(WzClient)
	wcc.Fqdn = payload[wzlib_transport.PAYLOAD_SYSTEM_FQDN].(string)
	wcc.Uid = payload[wzlib_transport.PAYLOAD_SYSTEM_ID].(string)
	wcc.RsaPk = string(payload[wzlib_transport.PAYLOAD_RSA].([]byte))

	return wcc.SetFingerprint()
}

// SetFingerprint calculates and sets the fingerprint from the RSA key
func (wcl *WzClient) SetFingerprint() *WzClient {
	if wcl.RsaPk == "" {
		panic("Setting fingerprint from an empty key!")
	}

	if wcl.RsaFp == "" {
		wcl.RsaFp = wzlib_crypto.NewWzCryptoUtils().PEMKeyFingerprintFromString(wcl.RsaPk)
	}
	return wcl
}
