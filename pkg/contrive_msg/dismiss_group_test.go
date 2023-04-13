package contrive_msg

import "testing"

func TestDismissGroup(t *testing.T) {
	type args struct {
		OperateID string
		UserID    string
		GroupID   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				OperateID: "123",
				UserID:    "1914080869",
				GroupID:   "670303005",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DismissGroup(tt.args.OperateID, tt.args.UserID, tt.args.GroupID); (err != nil) != tt.wantErr {
				t.Errorf("DismissGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 推送群聊抢红包消息
func TestRedPacketGrabPushToGroup(t *testing.T) {
	type args struct {
		OperateID        string
		SendPacketUserID string
		RedPacketID      string
		SendUserName     string
		ClickUserName    string
		GroupID          string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试群聊抢红包消息推送",
			args: args{
				OperateID:        "123",
				SendPacketUserID: "1914080869",
				RedPacketID:      "123",
				SendUserName:     "123",
				ClickUserName:    "123",
				GroupID:          "670303005",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RedPacketGrabPushToGroup(tt.args.OperateID, tt.args.SendPacketUserID, tt.args.RedPacketID, tt.args.SendUserName, tt.args.ClickUserName, tt.args.GroupID); (err != nil) != tt.wantErr {
				t.Errorf("RedPacketGrabPushToGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 测试给用户推送抢红包消息
func TestRedPacketGrabPushToUser(t *testing.T) {
	type args struct {
		OperateID        string
		SendPacketUserID string
		RedPacketID      string
		SendUserName     string
		ClickUserName    string
		GroupID          string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试给用户推送抢红包消息",
			args: args{
				OperateID:        "123",
				SendPacketUserID: "1914080869",
				RedPacketID:      "123",
				SendUserName:     "123",
				ClickUserName:    "123",
				GroupID:          "670303005",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RedPacketGrabPushToUser(tt.args.OperateID, tt.args.SendPacketUserID, tt.args.RedPacketID, tt.args.SendUserName, tt.args.ClickUserName, tt.args.GroupID); (err != nil) != tt.wantErr {
				t.Errorf("RedPacketGrabPushToUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendRedPacketLuckyMessage(t *testing.T) {
	type args struct {
		OperateID        string
		SendPacketUserID string
		RedPacketID      string
		LuckyUserName    string
		GroupID          string
		spendTime        int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试发送红包抢到消息",
			args: args{
				OperateID:        "123",
				SendPacketUserID: "1914080869",
				RedPacketID:      "123",
				LuckyUserName:    "123",
				GroupID:          "670303005",
				spendTime:        123,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendRedPacketLuckyMessage(tt.args.OperateID, tt.args.SendPacketUserID, tt.args.RedPacketID, tt.args.LuckyUserName, tt.args.GroupID, tt.args.spendTime); (err != nil) != tt.wantErr {
				t.Errorf("SendRedPacketLuckyMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
