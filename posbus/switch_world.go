package posbus

import (
	"github.com/momentum-xyz/posbus-protocol/utils"

	// External
	"github.com/google/uuid"
)

type SwitchWorld struct {
	*Message
}

func NewSwitchWorldMsg(id uuid.UUID) *SwitchWorld {
	obj := NewMessage(MsgTypeSwitchWorld, MsgUUIDTypeSize)
	copy(obj.Msg()[:MsgUUIDTypeSize], utils.BinID(id))
	return &SwitchWorld{
		Message: obj,
	}
}

func (m *Message) AsSwitchWorld() *SwitchWorld {
	return &SwitchWorld{
		Message: m,
	}
}

func (m *SwitchWorld) World() uuid.UUID {
	id, _ := uuid.FromBytes(m.Msg()[:MsgUUIDTypeSize])
	return id

}
