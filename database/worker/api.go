package wzlib_database_worker

import "github.com/jinzhu/gorm"

type WzWorkerDbh struct {
	clients *WzWorkerClientsAPI
	keys    *WzWorkerClientKeysAPI
	jobs    *WzWorkerJobAPI
}

func NewWzWorkerDbh() *WzWorkerDbh {
	wdb := new(WzWorkerDbh)
	wdb.clients = NewWzWorkerClientsAPI()
	wdb.keys = NewWzWorkerClientKeysAPI()
	wdb.jobs = NewWorkerJobAPI()
	return wdb
}

// Pass GORM database handler reference to the clients API
func (wdb *WzWorkerDbh) SetDbh(dbh *gorm.DB) {
	wdb.clients.setDbh(dbh)
	wdb.keys.setDbh(dbh)
	wdb.jobs.setDbh(dbh)
}

func (wdb *WzWorkerDbh) GetClientsAPI() *WzWorkerClientsAPI {
	if wdb.clients.db == nil {
		panic("Clients API database reference was not initalised")
	}
	return wdb.clients
}

func (wdb *WzWorkerDbh) GetKeysAPI() *WzWorkerClientKeysAPI {
	if wdb.keys.db == nil {
		panic("Keys API database reference was not initialised")
	}
	return wdb.keys
}

func (wdb *WzWorkerDbh) GetJobsAPI() *WzWorkerJobAPI {
	if wdb.jobs.db == nil {
		panic("Jobs API database reference was not initialised")
	}
	return wdb.jobs
}
