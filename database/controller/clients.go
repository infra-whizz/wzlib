package wzlib_database_controller

import (
	"github.com/infra-whizz/wzlib"
	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	"github.com/jinzhu/gorm"
)

// Everything with client
type WzCtrlClientsAPI struct {
	db *gorm.DB
	wzlib_logger.WzLogger
}

func NewWzCtrlClientsAPI() *WzCtrlClientsAPI {
	wcca := new(WzCtrlClientsAPI)
	return wcca
}

func (wcc *WzCtrlClientsAPI) setDbh(dbh *gorm.DB) *WzCtrlClientsAPI {
	wcc.db = dbh
	return wcc
}

// Register register a client that just appeared.
// Registration means "Your public RSA is in the database, now wait"
func (wcc *WzCtrlClientsAPI) Register(client *WzClient) int {
	var existing WzClient
	var status int
	wcc.db.Where("rsa_fp = ?", client.RsaFp).First(&existing)
	if existing.RsaFp == "" {
		client.Status = wzlib.CLIENT_STATUS_NEW
		wcc.db.Create(client)
		status = client.Status
		wcc.GetLogger().Infoln("Client", client.Fqdn, "has been registered")
	} else {
		wcc.GetLogger().Debugln("Client", client.Fqdn, "is already registered, skipping")
		status = existing.Status
	}
	return status
}

// Accept that was already registered.
// Accepetation means flipping status and it will be "OK, now you are in".
// This makes client listable for the workers.
// But the reconciliation needs to be extra called elsewhere.
func (wcc *WzCtrlClientsAPI) Accept() {}

// Reject sets its status as "rejected", but keeps in the database
// everything: FQDN, machine ID and RSA pubkey. Used for black-listing.
func (wcc *WzCtrlClientsAPI) Reject() {}

// Delete just deletes everything of the client.
// This client is eligible to be registered again.
func (wcc *WzCtrlClientsAPI) Delete() {}

// GetRegistered returns a list of new clients
func (wcc *WzCtrlClientsAPI) GetRegistered() []*WzClient {
	var clients []*WzClient
	wcc.db.Where("status = ?", wzlib.CLIENT_STATUS_NEW).Find(&clients)

	// Do not transfer RSA keys
	for _, system := range clients {
		system.RsaPk = ""
	}
	return clients
}

// GetRejected returns a list of new clients
func (wcc *WzCtrlClientsAPI) GetRejected() {}

// GetRegisteredAmount returns an amout of registered clients
func (wcc *WzCtrlClientsAPI) GetRegisteredAmount() int64 {
	return 0
}

// Search for clients based on specific query
func (wcc *WzCtrlClientsAPI) Search() {}

// GetByFQDN returns client data (struct?) by FQDN
func (wcc *WzCtrlClientsAPI) GetByFQDN() {}

// GetByUid returns client data (struct?) by system ID
func (wcc *WzCtrlClientsAPI) GetByUid() {}

// GetByFp returns client data (struct?) by RSA fingerprint
func (wcc *WzCtrlClientsAPI) GetByFp() {}

// Set sets/updates/adds client's status (struct?). This does not include traits.
func (wcc *WzCtrlClientsAPI) Set() {}

// UpdateTraits adds/sets/updates client's traits
func (wcc *WzCtrlClientsAPI) UpdateClientTraits() {}
