package wzlib_sockets

/*
Socket communicator is used to communicate locally between the instances.
It eliminates .pid files and management behind that. Typical use case is
to know if there is another instance already running or not.
*/

import (
	"fmt"
	"net"
	"os"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
)

// WzLocalSocketCommunicator is used to communicate between instances locally ovber Unix socket.
// Usually used to know if the instance is running, its status etc.
type WzLocalSocketCommunicator struct {
	name         string
	sockListener net.Listener
	sockConn     net.Conn
	wzlib_logger.WzLogger
}

// NewWzLocalSocketCommunicator creates new socket communicator instance
func NewWzLocalSocketCommunicator(name string) *WzLocalSocketCommunicator {
	wls := new(WzLocalSocketCommunicator)
	wls.name = name
	return wls
}

// Bind to the Unix socket
func (wls *WzLocalSocketCommunicator) Bind() error {
	_, err := os.Stat(wls.name)
	if os.IsNotExist(err) {
		wls.GetLogger().Debugf("Setting up an Unix socket at %s", wls.name)
		return wls.setServer()
	} else {
		if err = wls.setClient(); err != nil {
			wls.GetLogger().Debugf("Socket %s is there, but dead: %s Removing and setting up as a server.", wls.name, err.Error())
			if err = os.Remove(wls.name); err != nil {
				wls.GetLogger().Panicf("Unable to remove previous socket file: %s. Terminating.", err.Error())
			}
			return wls.setServer()
		}
	}

	return err
}

func (wls *WzLocalSocketCommunicator) setClient() error {
	var err error
	wls.sockConn, err = net.Dial("unix", wls.name)
	if err != nil {
		wls.GetLogger().Debugf("Unable to setup client for the Unix socket: %s", err.Error())
	}

	return err
}

func (wls *WzLocalSocketCommunicator) setServer() error {
	var err error
	wls.sockListener, err = net.Listen("unix", wls.name)
	if err != nil {
		wls.GetLogger().Debugf("Error setting up Unix socket server: %s", err.Error())
	}

	return err
}

// IsServer mode
func (wls *WzLocalSocketCommunicator) IsServer() bool {
	return wls.sockListener != nil
}

// IsClient mode
func (wls *WzLocalSocketCommunicator) IsClient() bool {
	return wls.sockConn != nil
}

func (wls *WzLocalSocketCommunicator) Init() error {
	return nil
}

// GetStatus from the server (client mode)
func (wls *WzLocalSocketCommunicator) GetStatus() error {
	if !wls.IsClient() {
		return fmt.Errorf("Socket is not in a client mode")
	}
	return nil
}

// Returns status about current instance (server mode)
func (wls *WzLocalSocketCommunicator) returnStatus() error {
	if !wls.IsServer() {
		return fmt.Errorf("Socket is not in a server mode")
	}
	return nil
}
