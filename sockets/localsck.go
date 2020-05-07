package wzlib_sockets

/*
Socket communicator is used to communicate locally between the instances.
It eliminates .pid files and management behind that. Typical use case is
to know if there is another instance already running or not.
*/

import (
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
)

// WzLocalSockMiddlewareItf interface
type WzLocalSockMiddlewareItf interface {
	Name() string
	OnCommand(command string) string
}

// WzLocalSocketCommunicator is used to communicate between instances locally ovber Unix socket.
// Usually used to know if the instance is running, its status etc.
type WzLocalSocketCommunicator struct {
	name         string
	sockListener net.Listener
	sockConn     net.Conn
	middlewares  []WzLocalSockMiddlewareItf
	wzlib_logger.WzLogger
}

// NewWzLocalSocketCommunicator creates new socket communicator instance
func NewWzLocalSocketCommunicator(name string) *WzLocalSocketCommunicator {
	wls := new(WzLocalSocketCommunicator)
	wls.name = name
	wls.middlewares = make([]WzLocalSockMiddlewareItf, 0)
	return wls
}

// RegisterMiddleware for processing socket communicator any custom protocol and actions
func (wls *WzLocalSocketCommunicator) RegisterMiddleware(middleware WzLocalSockMiddlewareItf) *WzLocalSocketCommunicator {
	wls.middlewares = append(wls.middlewares, middleware)
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

// Handle request. Limits to 512 bytes, enough for the commands.
func (wls *WzLocalSocketCommunicator) handleRequest(conn net.Conn) {
	buf := make([]byte, 0x200)
	offset, err := conn.Read(buf)
	if err != nil {
		wls.GetLogger().Errorf("Error reading input data while handling request: %s", err.Error())
		return
	}

	for _, mw := range wls.middlewares {
		resp := mw.OnCommand(strings.TrimSpace(string(buf[0:offset]))) + "\n"
		wls.GetLogger().Debugf("Middleware %s responds: %s", mw.Name(), string(resp))
		_, err = conn.Write([]byte(resp))
		if err != nil {
			wls.GetLogger().Errorf("Unable to respond on Unix socket request: %s", err.Error())
		}
	}
}

func (wls *WzLocalSocketCommunicator) serve() {
	wls.GetLogger().Debugf("Listening on Unix socket at %s", wls.name)
	for {
		conn, err := wls.sockListener.Accept()
		if err != nil {
			wls.GetLogger().Debugf("Unix socket: %s", err.Error())
			return
		} else {
			go wls.handleRequest(conn)
		}
	}
}

func (wls *WzLocalSocketCommunicator) setServer() error {
	var err error
	wls.sockListener, err = net.Listen("unix", wls.name)
	if err != nil {
		wls.GetLogger().Debugf("Error setting up Unix socket server: %s", err.Error())
		return err
	}

	// Cleanup socket on exit
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func(listener net.Listener, ch chan os.Signal) {
		sgn := <-ch
		wls.GetLogger().Infof("Interrupt: %s", sgn)
		wls.sockListener.Close()
		os.Exit(0)
	}(wls.sockListener, sigCh)

	go wls.serve()

	return nil
}

// IsServer mode
func (wls *WzLocalSocketCommunicator) IsServer() bool {
	return wls.sockListener != nil
}

// IsClient mode
func (wls *WzLocalSocketCommunicator) IsClient() bool {
	return wls.sockConn != nil
}

// Request the server (client mode)
func (wls *WzLocalSocketCommunicator) Request(command string) string {
	if !wls.IsClient() {
		return ""
	}

	repl := make([]byte, 0x400)
	wls.sockConn.Write([]byte(command))
	wls.sockConn.Read(repl)

	return string(repl)
}

// Teardown the socket communicator bind
func (wls *WzLocalSocketCommunicator) Teardown() {
	if wls.sockListener != nil {
		wls.sockListener.Close()
	}
	if wls.sockConn != nil {
		wls.sockConn.Close()
	}
}
