package wzlib_database_worker

import (
	"github.com/jinzhu/gorm"
)

///////////////////////////////////////////////////////
// Clients
type WzWorkerClientsAPI struct {
	db *gorm.DB
}

func NewWzWorkerClientsAPI() *WzWorkerClientsAPI {
	wca := new(WzWorkerClientsAPI)
	return wca
}

// Set GORM database handler
func (wca *WzWorkerClientsAPI) setDbh(db *gorm.DB) *WzWorkerClientsAPI {
	wca.db = db
	return wca
}

// GetAssignedClients returns a slice of assigned client systems to the current worker.
// If there is only one worker, then it will get all of the client systems at once.
func (wca *WzWorkerClientsAPI) GetAssignedClients() {}

///////////////////////////////////////////////////////
// Query clients
func (wca *WzWorkerClientsAPI) GetClientByMachineId(uid string) {}

// GetClientByFQDN returns client data and its traits by FQDN
func (wca *WzWorkerClientsAPI) GetClientByFQDN(fqdn string) {}

// GetClientByFf returns client data and its traits by RSA fingerprint
func (wca *WzWorkerClientsAPI) GetClientByFp(fp string) {}
