package posbus

import (
	"encoding/binary"
	"github.com/google/uuid"
)

type TriggerInteraction struct {
	*Message
}

func (m *Message) AsTriggerInteraction() *TriggerInteraction {
	return &TriggerInteraction{
		Message: m,
	}
}

func (m *TriggerInteraction) Kind() uint32 {
	return binary.LittleEndian.Uint32(m.Msg())
}

func (m *TriggerInteraction) Target() uuid.UUID {
	id, _ := uuid.FromBytes(m.Msg()[MsgTypeSize : MsgTypeSize+MsgUUIDTypeSize])
	return id
}

func (m *TriggerInteraction) Flag() int32 {
	return int32(binary.LittleEndian.Uint32(m.Msg()[MsgTypeSize+MsgUUIDTypeSize:]))
}

func (m *TriggerInteraction) Label() string {
	n := binary.LittleEndian.Uint32(m.Msg()[MsgTypeSize+MsgUUIDTypeSize+MsgTypeSize:])
	return string(m.Msg()[MsgTypeSize+MsgUUIDTypeSize+MsgTypeSize+MsgArrTypeSize : MsgTypeSize+MsgUUIDTypeSize+MsgTypeSize+MsgArrTypeSize+n])
}
