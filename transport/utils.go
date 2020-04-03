package wzlib_transport

import wzlib_logger "github.com/infra-whizz/wzlib/logger"

type WzEventMsgUtils struct {
	wzlib_logger.WzLogger
}

// NewWzEventMsgUtils creates a constructor for message utils
func NewWzEventMsgUtils() *WzEventMsgUtils {
	return new(WzEventMsgUtils)
}

// GetMessage converts message bytes into a WzGenericMessage type
func (mu *WzEventMsgUtils) GetMessage(data []byte) (envelope *WzGenericMessage) {
	mu.GetLogger().Debugln("Decoding message of", len(data), "bytes")
	envelope = NewWzGenericMessage()
	if err := envelope.LoadBytes(data); err != nil {
		mu.GetLogger().Errorln("Got garbled console message:", err.Error())
		envelope = nil
	}
	return
}
