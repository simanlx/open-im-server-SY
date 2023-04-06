package contrive_msg

import (
	"testing"
)

func TestSendRedMessage(t *testing.T) {
	req := &FPacket{
		PacketID:    "1111111",
		UserID:      "10086",
		PacketType:  1,
		IsLucky:     1,
		PacketTitle: "新年快乐",
		OperateID:   "123",
		RecvID:      "10081",
		IsExclusive: 0,
	}
	SendRedMessage(req, "123")
}
