package contrive_msg

import (
	pb "Open_IM/pkg/proto/cloud_wallet"
	pb2 "Open_IM/pkg/proto/push"
	open_im_sdk "Open_IM/pkg/proto/sdk_ws"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送红包消息给用户
func SendRedMessage(req *FPacket, redpacketID string) (*pb2.PushMsgResp, error) {
	// 创建红包消息
	p, err := NewPostMessage(req)
	// http post 发送红包消息
	url := "localhost:8080/v4/openim/sendmsg"
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return nil, nil
}

type FPacket struct {
	ID              int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;not null" json:"id"`
	PacketID        string `gorm:"column:packet_id;not null" json:"packet_id"`
	UserID          string `gorm:"column:user_id;not null" json:"user_id"`
	PacketType      int32  `gorm:"column:packet_type;not null" json:"packet_type"`
	IsLucky         int32  `gorm:"column:is_lucky;not null" json:"is_lucky"`
	ExclusiveUserID string `gorm:"column:exclusive_user_id;not null" json:"exclusive_user_id"`
	PacketTitle     string `gorm:"column:packet_title;not null" json:"packet_title"`
	Amount          int64  `gorm:"column:amount;not null" json:"amount"`
	Number          int32  `gorm:"column:number;not null" json:"number"`
	ExpireTime      int64  `gorm:"column:expire_time;not null" json:"expire_time"`
	MerOrderID      string `gorm:"column:mer_order_id;not null" json:"mer_order_id"`
	OperateID       string `gorm:"column:operate_id;not null" json:"operate_id"`
	RecvID          string `gorm:"column:recv_id;not null" json:"recv_id"`
	CreatedTime     int64  `gorm:"column:created_time;not null" json:"created_time"`
	UpdatedTime     int64  `gorm:"column:updated_time;not null" json:"updated_time"`
	Status          int32  `gorm:"column:status;not null" json:"status"` // 0 创建未生效，1 为红包正在领取中，2为红包领取完毕，3为红包过期
	IsExclusive     int32  `gorm:"column:is_exclusive;not null" json:"is_exclusive"`
}

/*senderPlatformID*/

type Data struct {
	SessionType int32                        `json:"sessionType" binding:"required"`
	MsgFrom     int32                        `json:"msgFrom" binding:"required"`
	ContentType int32                        `json:"contentType" binding:"required"`
	RecvID      string                       `json:"recvID" `
	GroupID     string                       `json:"groupID" `
	ForceList   []string                     `json:"forceList"`
	Content     ContriveData                 `json:"content" binding:"required"`
	Options     map[string]bool              `json:"options" `
	ClientMsgID string                       `json:"clientMsgID" binding:"required"`
	CreateTime  int64                        `json:"createTime" binding:"required"`
	OffLineInfo *open_im_sdk.OfflinePushInfo `json:"offlineInfo" `
}

type ContriveData struct {
	Data        string `json:"data"`
	Description string `json:"description"`
	Extension   string `json:"extension"`
}

// 发送红包消息
func NewPostMessage(f *FPacket) (*paramsUserSendMsg, error) {
	var (
		RecvID      string = ""
		sessionType int    = 0
		GroupID     string = ""
	)
	if f.PacketType == 1 {
		// 个人红包
		RecvID = f.RecvID
		sessionType = 1 // 单聊
	} else {
		// 群红包
		RecvID = ""
		GroupID = f.RecvID
		sessionType = 2 // 群消息
	}

	p := &paramsUserSendMsg{
		SenderPlatformID: 1,
		SendID:           f.UserID,
		OperationID:      f.OperateID,
		Data: Data{
			SessionType: int32(sessionType),
			MsgFrom:     1,
			ContentType: 110,
			RecvID:      RecvID,
			GroupID:     GroupID,
			ForceList:   []string{},
			Options:     map[string]bool{},
			ClientMsgID: f.OperateID,
			CreateTime:  time.Now().Unix(),
			OffLineInfo: &open_im_sdk.OfflinePushInfo{
				Title:        "你有新的红包",
				Desc:         "",
				Ex:           "",
				IOSPushSound: "default",
			},
		},
	}
	// 自定义红包消息

	res := NewRedPacket(f)
	content, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	data := ContriveData{
		Data:        string(content),
		Description: "红包消息",
		Extension:   "",
	}

	p.Data.Content = data
	return p, nil
}

// 创建一个红包消息
func NewRedPacket(f *FPacket) *pb.SendRedPacketReq {
	result := &pb.SendRedPacketReq{
		UserId:          f.UserID,          // 红包发起人
		PacketType:      f.PacketType,      // 红包类型 1.个人红包 2.群红包
		IsLucky:         f.IsLucky,         // 是否是拼手气红包
		IsExclusive:     f.IsExclusive,     // 是否是专属红包
		ExclusiveUserID: f.ExclusiveUserID, // 专属红包接收人
		PacketTitle:     f.PacketTitle,     // 红包标题
		RecvID:          f.RecvID,          // 红包接收人
	}
	return result
}

type GrabRedPacketReq struct {
	UserId   string `json:"userId"`   // 抢红包人
	PacketID string `json:"packetID"` // 红包ID
	Time     int64  `json:"time"`     // 抢红包时间
}
