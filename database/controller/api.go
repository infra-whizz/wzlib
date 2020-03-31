package wzlib_database_controller

import "github.com/jinzhu/gorm"

type WzControllerDbh struct {
	clients *WzCtrlClientsAPI
	workers *WzCtrlWorkersAPI
}

func NewWzControllerDbh() *WzControllerDbh {
	wcd := new(WzControllerDbh)
	wcd.clients = NewWzCtrlClientsAPI()
	wcd.workers = NewWzCtrlWorkersAPI()
	return wcd
}

func (wcd *WzControllerDbh) SetDbh(dbh *gorm.DB) {
	wcd.clients.setDbh(dbh)
	wcd.workers.setDbh(dbh)
}

func (wcd *WzControllerDbh) GetClientsAPI() *WzCtrlClientsAPI {
	if wcd.clients.db == nil {
		panic("Database is not set for Controller Clients API")
	}
	return wcd.clients
}
