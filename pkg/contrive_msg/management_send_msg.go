package contrive_msg

import (
	server_api_params "Open_IM/pkg/proto/sdk_ws"
	"encoding/json"
	"fmt"
)

type ManagementSendMsg struct {
	OperationID         string `json:"operationID" binding:"required"`
	BusinessOperationID string `json:"businessOperationID"`
	SendID              string `json:"sendID" binding:"required"`
	GroupID             string `json:"groupID" `
	SenderNickname      string `json:"senderNickname" `
	SenderFaceURL       string `json:"senderFaceURL" `
	SenderPlatformID    int32  `json:"senderPlatformID"`
	//ForceList        []string                     `json:"forceList" `
	Content         ContriveData                       `json:"content" binding:"required" swaggerignore:"true"`
	ContentType     int32                              `json:"contentType" binding:"required"`
	SessionType     int32                              `json:"sessionType" binding:"required"`
	IsOnlineOnly    bool                               `json:"isOnlineOnly"`
	NotOfflinePush  bool                               `json:"notOfflinePush"`
	OfflinePushInfo *server_api_params.OfflinePushInfo `json:"offlinePushInfo"`
}

func NewManagementSendMsg_RedMsg(f *FPacket) *ManagementSendMsg {

	res := &ManagementSendMsg{
		OperationID:         "1111111",
		BusinessOperationID: "111111111",
		SendID:              "10086",
		GroupID:             "10086",
		SenderPlatformID:    1,
		Content: ContriveData{
			Data:        "你好",
			Description: "ddddd",
			Extension:   "dddd",
		},
		ContentType:     110,
		SessionType:     1,
		IsOnlineOnly:    false,
		NotOfflinePush:  false,
		OfflinePushInfo: &server_api_params.OfflinePushInfo{},
	}

	co, _ := json.Marshal(res)
	fmt.Println(string(co))
	return nil
}
