package base_info

// 红包结构体 ： 推送红包结构体

type RedPacket struct {
	RedPacketId   string `json:"red_packet_id"`   // 红包id
	RedPacketName string `json:"red_packet_name"` // 红包名称
	RedPacketType string `json:"red_packet_type"` // 红包类型 1：普通红包 2：拼手气红包
	RedPacketNum  string `json:"red_packet_num"`  // 红包数量 1-200
	RedPacketAmt  string `json:"red_packet_amt"`  // 红包金额 单位：分
}
