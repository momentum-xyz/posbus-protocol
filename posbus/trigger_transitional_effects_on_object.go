package posbus

import (
	"encoding/binary"

	"github.com/momentum-xyz/posbus-protocol/utils"

	"github.com/google/uuid"
)

const TriggerTransitionalEffectsOnObjectElementSize = 2*MsgUUIDTypeSize + MsgTypeSize

type TriggerTransitionalEffectsOnObject struct {
	*Message
}

func NewTriggerTransitionalEffectsOnObjectMsg(numEffects int) *TriggerTransitionalEffectsOnObject {
	obj := NewMessage(MsgTypeTriggerTransitionalEffectsOnObject, MsgArrTypeSize+numEffects*TriggerTransitionalEffectsOnObjectElementSize)
	binary.LittleEndian.PutUint32(obj.Msg(), uint32(numEffects))
	return &TriggerTransitionalEffectsOnObject{
		Message: obj,
	}
}

func (m *TriggerTransitionalEffectsOnObject) SetEffect(i int, emitter, object uuid.UUID, effect uint32) {
	start := MsgArrTypeSize + i*TriggerTransitionalEffectsOnObjectElementSize
	copy(m.Msg()[start:], utils.BinID(emitter))
	copy(m.Msg()[start+MsgUUIDTypeSize:], utils.BinID(object))
	binary.LittleEndian.PutUint32(m.Msg()[start+2*MsgUUIDTypeSize:], effect)
}
