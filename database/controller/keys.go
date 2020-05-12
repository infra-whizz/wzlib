package wzlib_database_controller

import (
	"fmt"

	wzlib_crypto "github.com/infra-whizz/wzlib/crypto"
	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	"github.com/jinzhu/gorm"
)

// WzPEMKeyEntity entity object for the database of PEM keys
type WzPEMKeyEntity struct {
	ID        int    `gorm:"primary_key"`
	RsaFp     string `gorm:"unique; not null"`
	RsaPk     []byte `gorm:"unique; not null"`
	MachineId string `gorm:"unique;"`
	Fqdn      string `gorm:"unique;"`
}

type WzCtrlKeysAPI struct {
	db    *gorm.DB
	utils *wzlib_crypto.WzCryptoUtils
	wzlib_logger.WzLogger
}

func NewWzCtrlKeysAPI() *WzCtrlKeysAPI {
	wck := new(WzCtrlKeysAPI)
	wck.utils = wzlib_crypto.NewWzCryptoUtils()
	return wck
}

func (wck *WzCtrlKeysAPI) setDbh(dbh *gorm.DB) *WzCtrlKeysAPI {
	wck.db = dbh
	return wck
}

// AddRSAPublicPEM returns client's RSA public key in PEM format, queried by the machine ID.
// the fqdn is to merely indicate what machine is holding it, but the key is tied up to the machine ID.
func (wck *WzCtrlKeysAPI) AddRSAPublicPEM(keypem []byte, machineid string, fqdn string) error {
	if machineid == "" {
		return fmt.Errorf("Unable to add PEM key: machine ID required.")
	}
	var existing WzPEMKeyEntity
	fingerprint := wck.utils.PEMKeyFingerprintFromBytes(keypem)
	wck.db.Where("rsa_fp = ?", fingerprint).First(&existing)
	if existing.RsaFp == "" {
		wck.db.Create(&WzPEMKeyEntity{
			RsaFp:     fingerprint,
			RsaPk:     keypem,
			MachineId: machineid,
		})
	}

	return nil
}

// RemoveRSAPublicPEM from the database by full fingerprint
func (wck *WzCtrlKeysAPI) RemoveRSAPublicPEM(fingerprint string) error {
	if fingerprint == "" {
		return fmt.Errorf("Unable to remove PEM key: fingerprint required")
	}

	key := &WzPEMKeyEntity{}

	wck.db.Where("rsa_fp = ?", fingerprint).First(&key)
	if key.RsaFp == fingerprint {
		wck.db.Model(&key).Where("rsa_fp = ?", key.RsaFp).Delete(WzPEMKeyEntity{})
		wck.GetLogger().Infof("Deleted key for '%s' (fingerprint: %s...%s)", key.Fqdn, key.RsaFp[:8], key.RsaFp[len(key.RsaFp)-8:])
		return nil
	} else {
		return fmt.Errorf("PEM key not found by the given fingerprint")
	}
}