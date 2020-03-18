/*
	Author: Bo Maryniuk
	Bus connector
*/
package wzd_transport

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"strings"
)

type NatsURL struct {
	Scheme string
	Fqdn   string
	Port   int
}

type WzdPubSub struct {
	urls []*NatsURL
	ncp  *nats.Conn
	ncs  *nats.Conn
}

func NewWizPubSub() *WzdPubSub {
	wzd := new(WzdPubSub)
	wzd.urls = make([]*NatsURL, 0)
	return wzd
}

// AddNatsServerURL adds NATS server URL to the cluster of servers to connect
func (wzd *WzdPubSub) AddNatsServerURL(host string, port int) *WzdPubSub {
	wzd.urls = append(wzd.urls, &NatsURL{Scheme: "nats", Fqdn: host, Port: port})
	return wzd
}

// IsConnected currently only indicates if the connection is initialised
func (wzd *WzdPubSub) IsConnected() bool {
	return wzd.ncp != nil && wzd.ncs != nil
}

// Format cluster URLs
func (wzd *WzdPubSub) getClusterURLs() string {
	buff := make([]string, 0)
	for _, nurl := range wzd.urls {
		buff = append(buff, fmt.Sprintf("%s://%s:%d", nurl.Scheme, nurl.Fqdn, nurl.Port))
	}
	return strings.Join(buff, ", ")
}

// Connect to the cluster
func (wzd *WzdPubSub) connect() {
	var err error
	log.Printf("Connecting to %s...", wzd.getClusterURLs())
	if !wzd.IsConnected() {
		wzd.ncp, err = nats.Connect(wzd.getClusterURLs())
		log.Print("Connected publisher")
		if err != nil {
			log.Fatal(err)
		}
		wzd.ncs, err = nats.Connect(wzd.getClusterURLs())
		log.Print("Connected subscriber")
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Disconnect from the cluster
func (wzd *WzdPubSub) Disconnect() {
	if wzd.IsConnected() {
		log.Print("Begin disconnect")
		for _, nc := range [2]*nats.Conn{wzd.ncp, wzd.ncs} {
			if err := nc.Drain(); err != nil {
				log.Println(err.Error())
			}
			nc.Close()
		}
		wzd.ncp = nil
		wzd.ncs = nil
		log.Print("Disconected")
	}
}

func (wzd *WzdPubSub) GetPublisher() *nats.Conn {
	return wzd.ncp
}
func (wzd *WzdPubSub) GetSubscriber() *nats.Conn {
	return wzd.ncs
}

// Start starts the Node Controller
func (wzd *WzdPubSub) Start() {
	log.Print("Starting ncd event listener...")
	wzd.connect()
}
