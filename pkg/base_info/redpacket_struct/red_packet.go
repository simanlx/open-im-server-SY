package redpacket_struct

// 发送红包
type SendRedPacket struct {
	/*
		message SendRedPacketReq {
		  string userId = 1; //用户id
		  int32 PacketType = 2; //红包类型(1个人红包、2群红包)
		  int32 IsLucky = 3; //是否为拼手气红包
		  int32 IsExclusive = 4; //是否为专属红包(0不是、1是)
		  int32 ExclusiveUserID = 5; //专属红包接收者 和isExclusive
		  string PacketTitle = 6; //红包标题
		  int64 Amount = 7; //红包金额 单位：分
		  int32 Number = 8; //红包个数

		  // 通过哪种方式发送红包
		  int32 SendType = 9; //发送方式(1钱包余额、2银行卡)
		  int64 BankCardID = 10 ;//银行卡id
		  string operationID = 11; //链路跟踪id
		}
	*/
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
}

// 发送红包响应
type SendRedPacketResp struct {
}
