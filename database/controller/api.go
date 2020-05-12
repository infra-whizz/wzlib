package wzlib_database_controller

import "github.com/jinzhu/gorm"

// WzControllerDbh class bundle
type WzControllerDbh struct {
	clients *WzCtrlClientsAPI
	workers *WzCtrlWorkersAPI
	keys    *WzCtrlKeysAPI
}

// NewWzControllerDbh creates a new instances of the WzControllerDbh
func NewWzControllerDbh() *WzControllerDbh {
	wcd := new(WzControllerDbh)
	wcd.clients = NewWzCtrlClientsAPI()
	wcd.workers = NewWzCtrlWorkersAPI()
	wcd.keys = NewWzCtrlKeysAPI()
	return wcd
}

// SetDbh sets database handler to all the sub-objects
func (wcd *WzControllerDbh) SetDbh(dbh *gorm.DB) {
	wcd.clients.setDbh(dbh)
	wcd.workers.setDbh(dbh)
	wcd.keys.setDbh(dbh)
}

// GetClientsAPI returns an API bundle to access Client facility
func (wcd *WzControllerDbh) GetClientsAPI() *WzCtrlClientsAPI {
	if wcd.clients.db == nil {
		panic("Database is not set for Controller Clients API")
	}
	return wcd.clients
}

// GetKeysAPI returns an API bundle to store/retrieve/search for the PEM keys
func (wcd *WzControllerDbh) GetKeysAPI() *WzCtrlKeysAPI {
	if wcd.keys.db == nil {
		panic("Database is not set for Controller PubKeys infrastructure API")
	}
	return wcd.keys
}
