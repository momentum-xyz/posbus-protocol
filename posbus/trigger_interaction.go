package posbus

import (
	"encoding/binary"
	"github.com/google/uuid"
	"github.com/momentum-xyz/posbus-protocol/utils"
)

type TriggerInteraction struct {
	*Message
}

func NewTriggerInteractionMsg(kind uint32, target uuid.UUID, flag int32, label string) *TriggerInteraction {
	obj := NewMessage(MsgTriggerInteraction, MsgTypeSize+MsgUUIDTypeSize+MsgTypeSize+MsgArrTypeSize+len(label))
	binary.LittleEndian.PutUint32(obj.Msg(), kind)
	copy(obj.Msg()[MsgTypeSize:], utils.BinID(target))
	binary.LittleEndian.PutUint32(obj.Msg()[MsgTypeSize+MsgUUIDTypeSize:], uint32(flag))
	binary.LittleEndian.PutUint32(obj.Msg()[MsgTypeSize+MsgUUIDTypeSize+MsgTypeSize:], uint32(len(label)))
	copy(obj.Msg()[MsgTypeSize+MsgUUIDTypeSize+MsgTypeSize+MsgArrTypeSize:], label)
	return &TriggerInteraction{
		Message: obj,
	}
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
