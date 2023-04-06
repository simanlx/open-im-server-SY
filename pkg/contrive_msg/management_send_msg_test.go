package contrive_msg

import (
	"testing"
)

func TestNewManagementSendMsg_RedMsg(t *testing.T) {
	req := &FPacket{
		PacketID:    "1111111",
		UserID:      "1914080869",
		PacketType:  1,
		IsLucky:     1,
		PacketTitle: "新年快乐",
		OperateID:   "123",
		RecvID:      "1914080869",
		IsExclusive: 0,
	}
	SendSendRedPacket(req, "10000085")
}

//
func TestSendSendRedPacket(t *testing.T) {
	SendGrabPacket("1914080869", "1914080869", 1, "123", "100", "", "1111")
}
