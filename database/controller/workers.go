package wzlib_database_controller

import (
	"github.com/jinzhu/gorm"
)

type WzCtrlWorkersAPI struct {
	db *gorm.DB
}

func NewWzCtrlWorkersAPI() *WzCtrlWorkersAPI {
	cwa := new(WzCtrlWorkersAPI)
	return cwa
}

func (cwa *WzCtrlWorkersAPI) setDbh(dbh *gorm.DB) *WzCtrlWorkersAPI {
	cwa.db = dbh
	return cwa
}

// GetAll workers
func (cwa *WzCtrlWorkersAPI) GetAll() {}

// GetAssignedClients per a worker
func (cwa *WzCtrlWorkersAPI) GetAssignedClients() {
	/*
		this should accept a worker's ID and return
		a list of clients, assigned. to it.
	*/
}

// AssignClients will assign clients to the worker
func (cwa *WzCtrlWorkersAPI) AssignClients() {
	/*
		Accept an array of client IDs
	*/
}

// UnassignClients from the worker
func (cwa *WzCtrlWorkersAPI) UnassignClients() {
	/*
		Accept an array of client IDs
	*/
}

// Remove worker from the cluster (unassign everything)
func (cwa *WzCtrlWorkersAPI) Remove() {}

// Add worker to the cluster (keeps empty, needs reconciliation)
func (cwa *WzCtrlWorkersAPI) Add() {}

// Reconcile the whole cluster. This gets to each worker and
// reassigns clients to work with. Each worker works only
// with assigned clients and every time there might be a new
// clients, it never met before. Worker, essentially, a job proxy.
//
// The reconsiliation algorithm is to shift as little as possible
// of clients, so only necessary machines are re-assigned to another
// worker.
func (cwa *WzCtrlWorkersAPI) Reconcile() {}

// Activate a worker. This should return a reasonable timeout
// after which Reconcile() is called. If Activate is called again
// (new worker is in), then the previous timeout for the entire
// reconciliation is updated with this one
func (cwa *WzCtrlWorkersAPI) Activate() {}

// Deactivate a worker (reconciles the cluster)
func (cwa *WzCtrlWorkersAPI) Deactivate() {}
