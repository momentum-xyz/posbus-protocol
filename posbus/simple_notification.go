package posbus

import (
	"encoding/binary"
)

type SimpleNotification struct {
	*Message
}

func NewSimpleNotificationMsg(dest Destination, kind Notification, flag int32, s string) *SimpleNotification {
	obj := NewMessage(MsgTypeSimpleNotification, 1+2*MsgTypeSize+MsgArrTypeSize+len(s))
	obj.Msg()[0] = byte(dest)
	binary.LittleEndian.PutUint32(obj.Msg()[1:], uint32(kind))
	binary.LittleEndian.PutUint32(obj.Msg()[1+MsgTypeSize:], uint32(flag))
	binary.LittleEndian.PutUint32(obj.Msg()[1+2*MsgTypeSize:], uint32(len(s)))
	copy(obj.Msg()[1+2*MsgTypeSize+MsgArrTypeSize:], s)
	return &SimpleNotification{
		Message: obj,
	}
}
