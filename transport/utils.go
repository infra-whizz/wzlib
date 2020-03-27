package wzlib_transport

import (
	"log"
)

type WzEventMsgUtils struct{}

// NewWzEventMsgUtils creates a constructor for message utils
func NewWzEventMsgUtils() *WzEventMsgUtils {
	return new(WzEventMsgUtils)
}

// GetMessage converts message bytes into a WzGenericMessage type
func (mu *WzEventMsgUtils) GetMessage(data []byte) (envelope *WzGenericMessage) {
	log.Println("Decoding message of", len(data), "bytes")
	envelope = NewWzGenericMessage()
	if err := envelope.LoadBytes(data); err != nil {
		log.Println("Got garbled console message:", err.Error())
		envelope = nil
	}
	return
}
