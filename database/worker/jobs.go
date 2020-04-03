package wzlib_database_worker

import (
	"time"

	"github.com/jinzhu/gorm"
)

///////////////////////////////////////////////////////
// Job
type WzWorkerJobAPI struct {
	db *gorm.DB
}

func NewWorkerJobAPI() *WzWorkerJobAPI {
	wj := new(WzWorkerJobAPI)
	return wj
}

func (wdb *WzWorkerJobAPI) setDbh(dbh *gorm.DB) *WzWorkerJobAPI {
	wdb.db = dbh
	return wdb
}

// GetAllByClient returns all jobs, known in the log by the client
func (wdb *WzWorkerJobAPI) GetAllByClient(client string, till *time.Time) {}

func (wdb *WzWorkerJobAPI) GetUnfinished(client string) {}

func (wdb *WzWorkerJobAPI) GetFailed(client string) {}

func (wdb *WzWorkerJobAPI) New(client string) {}

func (wdb *WzWorkerJobAPI) FlushAfter(point *time.Time) {}

func (wdb *WzWorkerJobAPI) Get(jid string) {}

func (wdb *WzWorkerJobAPI) Delete(jid string) {}

func (wdb *WzWorkerJobAPI) Update(jid string) {}
