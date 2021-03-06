/*
	Author: Bo Maryniuk
	Bus connector
*/
package wzlib_transport

import (
	"fmt"
	"strings"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	"github.com/nats-io/nats.go"
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

	wzlib_logger.WzLogger
}

func NewWizPubSub() *WzdPubSub {
	wzd := new(WzdPubSub)
	wzd.urls = make([]*NatsURL, 0)
	return wzd
}

// AddNatsServerURL adds NATS server URL to the cluster of servers to connect
func (wzd *WzdPubSub) AddNatsServerURL(host string, port int) *WzdPubSub {
	wzd.GetLogger().Printf("Registering bus at %s:%d", host, port)
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
	wzd.GetLogger().Infof("Connecting to %s...", wzd.getClusterURLs())
	if !wzd.IsConnected() {
		wzd.ncp, err = nats.Connect(wzd.getClusterURLs())
		wzd.GetLogger().Infoln("Connected publisher")
		if err != nil {
			wzd.GetLogger().Fatal(err)
		}
		wzd.ncs, err = nats.Connect(wzd.getClusterURLs())
		wzd.GetLogger().Infoln("Connected subscriber")
		if err != nil {
			wzd.GetLogger().Fatal(err)
		}
	}
}

// Disconnect from the cluster
func (wzd *WzdPubSub) Disconnect() {
	if wzd.IsConnected() {
		wzd.GetLogger().Debugln("Begin disconnect")
		for _, nc := range [2]*nats.Conn{wzd.ncp, wzd.ncs} {
			if err := nc.Drain(); err != nil {
				wzd.GetLogger().Errorln(err.Error())
			}
			nc.Close()
		}
		wzd.ncp = nil
		wzd.ncs = nil
		wzd.GetLogger().Infoln("Disconected")
	}
}

func (wzd *WzdPubSub) GetPublisher() *nats.Conn {
	return wzd.ncp
}

func (wzd *WzdPubSub) PublishEnvelopeToChannel(channel string, envelope *WzGenericMessage) {
	data, err := envelope.Serialise()
	if err != nil {
		wzd.GetLogger().Errorln("Error serialising envelope:", err.Error())
	} else {
		if err := wzd.GetPublisher().Publish(channel, data); err != nil {
			wzd.GetLogger().Errorln("Error publishing message:", err.Error())
		}
	}
}

func (wzd *WzdPubSub) GetSubscriber() *nats.Conn {
	return wzd.ncs
}

// Start starts the Node Controller
func (wzd *WzdPubSub) Start() {
	wzd.GetLogger().Infoln("Starting ncd event listener...")
	wzd.connect()
}
