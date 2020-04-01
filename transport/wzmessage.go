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

// Message keys
const (
	PAYLOAD_RSA         = "rsa.pub"     // cipher text in PEM
	PAYLOAD_SYSTEM_ID   = "system.id"   // unique string of system ID
	PAYLOAD_SYSTEM_FQDN = "system.fqdn" // Host FQDN or just a hostname
	PAYLOAD_PING_ID     = "ping.id"     // ID of a ping request

	/*
		Function return payload. The value is a nested string/interface mapping
		which can contain whatever specific. This resides one level deeper from
		the fixed keys.
	*/
	PAYLOAD_FUNC_RET = "function.return"

	/*
		Size of batch. If a message is way too big, it should be then splitted
		into a series of those. Each batch message should always contain the
		same JID, so they then can be rejoined on the other hand back into
		one message.

		The "batch.size" denotes N messages: "X of N".
	*/
	PAYLOAD_BATCH_SIZE = "batch.size"

	PAYLOAD_COMMAND = "command" // specific key of the command, still ad-hoc :-(
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
