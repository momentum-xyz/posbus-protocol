package posbus

import (
	"encoding/binary"
)

type SignalMessage struct {
	*Message
}

func (m *Message) AsSignal() *SwitchWorld {
	return &SwitchWorld{
		Message: m,
	}
}

func (m *SwitchWorld) Signal() Signal {
	return Signal(binary.LittleEndian.Uint32(m.Msg()[:MsgTypeSize]))
}

func NewSignalMsg(signal Signal) *SignalMessage {
	obj := NewMessage(MsgTypeSignal, MsgTypeSize)
	binary.LittleEndian.PutUint32(obj.Msg(), uint32(signal))
	return &SignalMessage{
		Message: obj,
	}
}
