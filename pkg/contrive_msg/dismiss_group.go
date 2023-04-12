package contrive_msg

import (
	server_api_params "Open_IM/pkg/proto/sdk_ws"
	"encoding/json"
)

// 解散群聊
func DismissGroup(OperateID, UserID, GroupID string) error {
	GroupDismissMsg := &ContriveMessage{
		Data:    &GroupDismissMessage{GroupID: GroupID},
		MsgType: 11,
	}
	co, _ := json.Marshal(GroupDismissMsg)
	res := &ManagementSendMsg{
		OperationID:         OperateID,
		BusinessOperationID: OperateID,
		SendID:              UserID,
		SenderPlatformID:    1,
		Content: ContriveData{
			Data:        string(co),
			Description: "解散群聊消息",
			Extension:   "",
		},
		ContentType:     110, // 自定义消息
		SessionType:     2,   // 1 单聊 2 群聊
		IsOnlineOnly:    false,
		NotOfflinePush:  false,
		GroupID:         GroupID, // 接收方ID 群聊
		OfflinePushInfo: &server_api_params.OfflinePushInfo{},
	}

	coo, _ := json.Marshal(res)
	return SendMessage(OperateID, coo)
}
