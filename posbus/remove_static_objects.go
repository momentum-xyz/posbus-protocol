package posbus

import (
	"encoding/binary"

	"github.com/momentum-xyz/posbus-protocol/utils"

	"github.com/google/uuid"
)

const RemoveStaticObjectsElementSize = MsgUUIDTypeSize

type RemoveStaticObjects struct {
	*Message
}

func NewRemoveStaticObjectsMsg(numObjects int) *RemoveStaticObjects {
	obj := NewMessage(MsgTypeRemoveStaticObjects, MsgArrTypeSize+numObjects*RemoveStaticObjectsElementSize)
	binary.LittleEndian.PutUint32(obj.Msg(), uint32(numObjects))
	return &RemoveStaticObjects{
		Message: obj,
	}
}

func (m *RemoveStaticObjects) SetObject(i int, id uuid.UUID) {
	start := MsgArrTypeSize + i*RemoveStaticObjectsElementSize
	copy(m.Msg()[start:], utils.BinID(id))
}
