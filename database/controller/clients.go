package wzlib_database_controller

import (
	"regexp"
	"strings"

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

func (wcc *WzCtrlClientsAPI) removeClientsRSA(clients []*WzClient) {
	// Do not transfer RSA keys
	for _, system := range clients {
		system.RsaPk = ""
	}
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
func (wcc *WzCtrlClientsAPI) Accept(fingerprints ...string) (missing []string) {
	// XXX: This should also return a list of accepted, so the target machines will get a notification
	return wcc.setStatusByFingerprints(wzlib.CLIENT_STATUS_ACCEPTED, "Accepted", fingerprints...)
}

// Reject sets its status as "rejected", but keeps in the database
// everything: FQDN, machine ID and RSA pubkey. Used for black-listing.
func (wcc *WzCtrlClientsAPI) Reject(fingerprints ...string) (missing []string) {
	// XXX: This should also return a list of rejected, so the target machines will get a notification
	return wcc.setStatusByFingerprints(wzlib.CLIENT_STATUS_REJECTED, "Rejected", fingerprints...)
}

func (wcc *WzCtrlClientsAPI) setStatusByFingerprints(status int, msg string, fingerprints ...string) (missing []string) {
	missing = make([]string, 0)
	if len(fingerprints) == 0 {
		// all at once
		client := &WzClient{}
		wcc.db.Model(&client).Where("status = ?", wzlib.CLIENT_STATUS_NEW).Update("status", status)
	} else {
		// by fingerprints
		for _, fp := range fingerprints {
			client := &WzClient{}
			wcc.db.Where("status = ? AND rsa_fp LIKE ?", wzlib.CLIENT_STATUS_NEW, fp+"%").First(&client)
			if client.RsaFp != "" {
				finger := client.RsaFp
				wcc.db.Model(&client).Where("status = ? AND rsa_fp = ?", wzlib.CLIENT_STATUS_NEW, finger).Update("status", status)
				wcc.GetLogger().Infof("%s '%s' (key: %s...%s)", msg, client.Fqdn, client.RsaFp[:8], client.RsaFp[len(client.RsaFp)-8:])
			} else {
				missing = append(missing, fp)
			}
		}
	}
	return
}

// Delete just deletes everything of the client.
// This client is eligible to be registered again.
func (wcc *WzCtrlClientsAPI) Delete(fingerprints ...string) (missing []string) {
	missing = make([]string, 0)
	for _, fp := range fingerprints {
		client := &WzClient{}
		wcc.db.Where("rsa_fp LIKE ?", fp+"%").First(&client)
		if client.RsaFp != "" {
			wcc.db.Model(&client).Where("rsa_fp = ?", client.RsaFp).Delete(WzClient{})
			wcc.GetLogger().Infof("Deleted '%s' (key: %s...%s)", client.Fqdn, client.RsaFp[:8], client.RsaFp[len(client.RsaFp)-8:])
		} else {
			missing = append(missing, fp)
		}
	}

	// XXX: This should also return a list of deleted, so the target machines will get a notification
	return
}

// GetRegistered returns a list of new clients
func (wcc *WzCtrlClientsAPI) GetRegistered() []*WzClient {
	var clients []*WzClient
	wcc.db.Where("status = ?", wzlib.CLIENT_STATUS_NEW).Find(&clients)
	wcc.removeClientsRSA(clients)
	return clients
}

// GetRejected returns a list of new clients
func (wcc *WzCtrlClientsAPI) GetRejected() []*WzClient {
	var clients []*WzClient
	wcc.db.Where("status = ?", wzlib.CLIENT_STATUS_REJECTED).Find(&clients)
	wcc.removeClientsRSA(clients)
	return clients
}

// GetRegisteredAmount returns an amout of registered clients
func (wcc *WzCtrlClientsAPI) GetRegisteredAmount() int64 {
	return 0
}

// Search for clients based on specific query
func (wcc *WzCtrlClientsAPI) Search(query string) []*WzClient {
	var clients []*WzClient

	if strings.Contains(query, "*") {
		wcc.db.Where("fqdn LIKE ?", regexp.MustCompile(`\*+`).ReplaceAllString(query, "%")).Find(&clients)
	} else {
		wcc.db.Where("fqdn = ?", query).Find(&clients)
	}
	wcc.removeClientsRSA(clients)
	return clients
}

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
