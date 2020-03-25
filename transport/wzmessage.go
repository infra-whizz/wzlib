package wzlib_transport

import (
	"github.com/infra-whizz/wzlib"
	"github.com/vmihailenco/msgpack/v4"
)

const (
	MSGTYPE_REGISTRATION = iota
	MSGTYPE_RUN_RESULT
)

// Console Message
type WzGenericMessage struct {
	Jid     string
	Type    int
	Payload map[string]interface{}
}

// Constructor
func NewWzGenericMessage(msgType int) *WzGenericMessage {
	wcm := new(WzGenericMessage)
	wcm.Jid = wzlib.MakeJid()
	wcm.Payload = make(map[string]interface{})
	wcm.Type = msgType

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
