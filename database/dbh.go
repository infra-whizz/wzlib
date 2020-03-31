/*
	Database API for the backend.
	Used by controllers and workers to track client systems,
	as well as the whole cluster information.
*/

package wzlib_database

import (
	"fmt"
	"log"

	wzlib_database_controller "github.com/infra-whizz/wzlib/database/controller"
	wzlib_database_worker "github.com/infra-whizz/wzlib/database/worker"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type WzDBH struct {
	controller *wzlib_database_controller.WzControllerDbh
	worker     *wzlib_database_worker.WzWorkerDbh

	// Db conn
	_dbUser string
	_dbHost string
	_dbPort int
	_dbName string
	_db     *gorm.DB

	// SSL
	_sslRootCert string
	_sslKey      string
	_sslCert     string
}

func NewWzDBH() *WzDBH {
	dbh := new(WzDBH)
	dbh._dbHost = "localhost"
	dbh._dbPort = 26257
	return dbh
}

// SetHost sets a host to connect
func (dbh *WzDBH) SetHost(host string) *WzDBH {
	dbh._dbHost = host
	return dbh
}

// SetPort sets a port to connect
func (dbh *WzDBH) SetPort(port int) *WzDBH {
	dbh._dbPort = port
	return dbh
}

// SetDbName sets a name of the database to connect
func (dbh *WzDBH) SetDbName(name string) *WzDBH {
	dbh._dbName = name
	return dbh
}

// SetSSLConf puts the ssl configuration
func (dbh *WzDBH) SetSSLConf(rootcert string, key string, sslcert string) *WzDBH {
	dbh._sslRootCert = rootcert
	dbh._sslKey = key
	dbh._sslCert = sslcert
	return dbh
}

// SetUser sets user of the database
func (dbh *WzDBH) SetUser(user string) *WzDBH {
	dbh._dbUser = user
	return dbh
}

// Open database connection
func (dbh *WzDBH) Open() {
	log.Println("Connecting to the database")
	var err error
	url := fmt.Sprintf("postgresql://%s@%s:%d/%s?ssl=true&sslmode=require&sslrootcert=%s&sslkey=%s&sslcert=%s",
		dbh._dbUser, dbh._dbHost, dbh._dbPort, dbh._dbName, dbh._sslRootCert, dbh._sslKey, dbh._sslCert)

	dbh._db, err = gorm.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database")
	dbh.automigrate()

	if dbh.controller != nil {
		dbh.controller.SetDbh(dbh._db)
	} else if dbh.worker != nil {
		dbh.worker.SetDbh(dbh._db)
	}
}

// Create or update all existing tables
func (dbh *WzDBH) automigrate() {
	dbh._db.AutoMigrate(&wzlib_database_controller.WzClient{})
}

// Close database connection
func (dbh *WzDBH) Close() {
	if dbh._db != nil {
		dbh._db.Close()
	} else {
		log.Println("Attempting to close unopened database reference")
	}
}

// WithWorkerAPI will initialise worker API and will nil controller API
func (dbh *WzDBH) WithWorkerAPI() *WzDBH {
	dbh.controller = nil
	dbh.worker = wzlib_database_worker.NewWzWorkerDbh()
	return dbh
}

// WithControllerAPI will initialise controller API and will nil worker API
func (dbh *WzDBH) WithControllerAPI() *WzDBH {
	dbh.worker = nil
	dbh.controller = wzlib_database_controller.NewWzControllerDbh()
	return dbh
}

// GetControllerAPI returns API for controller's interaction
func (dbh *WzDBH) GetControllerAPI() *wzlib_database_controller.WzControllerDbh {
	return dbh.controller
}

// GetWorkerAPI returns API for worker's interaction
func (dbh *WzDBH) GetWorkerAPI() *wzlib_database_worker.WzWorkerDbh {
	return dbh.worker
}
