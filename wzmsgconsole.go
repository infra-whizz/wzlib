package wzlib

import (
	"github.com/vmihailenco/msgpack/v4"
)

// Console Message
type WzConsoleMessage struct {
	Jid     string
	Query   string
	Command string
	kwargs  map[string]interface{}
}

// Constructor
func NewWzConsoleMessage() *WzConsoleMessage {
	wcm := new(WzConsoleMessage)
	wcm.Jid = MakeJid()
	wcm.kwargs = make(map[string]interface{})

	return wcm
}

// Serialise to bytes array
func (wcm *WzConsoleMessage) Serialise() ([]byte, error) {
	return msgpack.Marshal(wcm)
}

// LoadBytes loads the message from bytes
func (wcm *WzConsoleMessage) LoadBytes(data []byte) error {
	return msgpack.Unmarshal(data, wcm)
}

// Reply to the Console Message
type WzConsoleReplyMessage struct {
	Jid     string
	Msg     string // Error message
	Errcode int
}

func NewWzConsoleReplyMessage() *WzConsoleReplyMessage {
	return new(WzConsoleReplyMessage)
}

// Serialise to bytes array
func (wcrm *WzConsoleReplyMessage) Serialise() ([]byte, error) {
	return msgpack.Marshal(wcrm)
}

// LoadBytes loads the message from bytes
func (wcrm *WzConsoleReplyMessage) LoadBytes(data []byte) error {
	return msgpack.Unmarshal(data, wcrm)
}
