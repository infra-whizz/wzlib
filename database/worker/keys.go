package wzlib_database_worker

import (
	"github.com/jinzhu/gorm"
)

///////////////////////////////////////////////////////
// Client keys
type WzWorkerClientKeysAPI struct {
	db *gorm.DB
}

func NewWzWorkerClientKeysAPI() *WzWorkerClientKeysAPI {
	wck := new(WzWorkerClientKeysAPI)
	return wck
}

func (wdb *WzWorkerClientKeysAPI) setDbh(dbh *gorm.DB) *WzWorkerClientKeysAPI {
	wdb.db = dbh
	return wdb
}

// GetRSAPubkeyByMachineId returns client's RSA public key in PEM format, queried by the machine ID.
func (wdb *WzWorkerClientKeysAPI) GetRSAPubkeyByMachineId(uid string) {}

// GetRSAPubkeyByFQDN returns client's RSA public key in PEM format, queried by the FQDN.
func (wdb *WzWorkerClientKeysAPI) GetRSAPubkeyByFQDN(fqdn string) {}

// GetRSAPubkeyByFp returns client's RSA public key in PEM format, queried by the key fingerprint.
func (wdb *WzWorkerClientKeysAPI) GetRSAPubkeyByFp(fp string) {}
