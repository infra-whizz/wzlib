package wzlib_transport

import (
	"github.com/infra-whizz/wzlib"
	"github.com/vmihailenco/msgpack/v4"
)

const (
	MSGTYPE_GENERIC = iota // Also undefined when loading
	MSGTYPE_REGISTRATION
	MSGTYPE_PING
	MSGTYPE_RUN_RESULT
	MSGTYPE_CLIENT
)

// Console Message
type WzGenericMessage struct {
	Jid     string
	Type    int
	Payload map[string]interface{}
}

// NewWzMessage creates a message of a given type with Jid
func NewWzMessage(msgType int) *WzGenericMessage {
	wcm := new(WzGenericMessage)
	wcm.Jid = wzlib.MakeJid()
	wcm.Payload = make(map[string]interface{})
	wcm.Type = msgType

	return wcm
}

// NewWzGenericMessage creates a generic type of a message with no Jid
func NewWzGenericMessage() *WzGenericMessage {
	wcm := new(WzGenericMessage)
	wcm.Payload = make(map[string]interface{})
	wcm.Type = MSGTYPE_GENERIC

	return wcm
}

// Serialise to bytes array
func (wcm *WzGenericMessage) Serialise() ([]byte, error) {
	return msgpack.Marshal(wcm)
}

// LoadBytes loads the message from bytes
func (wcm *WzGenericMessage) LoadBytes(data []byte) error {
	return msgpack.Unmarshal(data, wcm)
}
