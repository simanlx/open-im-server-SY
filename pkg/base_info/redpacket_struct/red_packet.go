package redpacket_struct

// 发送红包
type SendRedPacket struct {
	UserId          string `json:"userId"`          //用户id
	PacketType      int32  `json:"PacketType"`      //红包类型(1个人红包、2群红包)
	IsLucky         int32  `json:"IsLucky"`         //是否为拼手气红包
	IsExclusive     int32  `json:"IsExclusive"`     //是否为专属红包(0不是、1是)
	ExclusiveUserID string `json:"ExclusiveUserID"` //专属红包接收者 和isExclusive
	PacketTitle     string `json:"PacketTitle"`     //红包标题
	Amount          int64  `json:"Amount"`          //红包金额 单位：分
	Number          int32  `json:"Number"`          //红包个数
	SendType        int32  `json:"SendType"`        //发送方式(1钱包余额、2银行卡)
	BankCardID      int64  `json:"BankCardID"`      //银行卡id
	OperationID     string `json:"operationID"`     //链路跟踪id
	RecvID          string `json:"recvID"`          //接收者id

	// 	BindCardAgrNo string `protobuf:"bytes,13,opt,name=bindCardAgrNo,proto3" json:"bindCardAgrNo,omitempty"` //绑卡协议号
	BindCardAgrNo string `json:"bindCardAgrNo"` //绑卡协议号
}

// 发送红包响应
type SendRedPacketResp struct {
}

// 点击抢红包
type ClickRedPacketReq struct {
	RedPacketID string `json:"redPacketID"` //红包id
	OperateID   string `json:"operateID"`   //链路跟踪id
}