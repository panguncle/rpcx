package protocol

import "sync"

/**
PReader: 使用了大量的sync.Pool来进行优化
https://juejin.im/post/5d4087276fb9a06adb7fbe4a
 */
var msgPool = sync.Pool{
	New: func() interface{} {
		header := Header([12]byte{})
		header[0] = magicNumber

		return &Message{
			Header: &header,
		}
	},
}

// GetPooledMsg gets a pooled message.
func GetPooledMsg() *Message {
	return msgPool.Get().(*Message)
}

// FreeMsg puts a msg into the pool.
func FreeMsg(msg *Message) {
	if msg != nil {
		msg.Reset()
		msgPool.Put(msg)
	}
}

var poolUint32Data = sync.Pool{
	New: func() interface{} {
		data := make([]byte, 4)
		return &data
	},
}
