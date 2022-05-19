package posbus

import "encoding/binary"

type RelayToReact struct {
	*Message
}

func NewRelayToReactMsg(module string, message []byte) *RelayToReact {
	obj := NewMessage(MsgTypeRelayToReact, MsgArrTypeSize+len(module)+MsgArrTypeSize+len(message))
	binary.LittleEndian.PutUint32(obj.Msg(), uint32(len(module)))
	copy(obj.Msg()[MsgArrTypeSize:], module)
	binary.LittleEndian.PutUint32(obj.Msg()[MsgArrTypeSize+len(module):], uint32(len(message)))
	copy(obj.Msg()[MsgArrTypeSize+len(module)+MsgArrTypeSize:], message)
	return &RelayToReact{
		Message: obj,
	}
}
