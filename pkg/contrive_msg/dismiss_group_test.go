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
			name: "手气红包",
			args: args{
				OperateID:        "123",
				SendPacketUserID: "1914080869",
				RedPacketID:      "123",
				LuckyUserName:    "steven2",
				GroupID:          "3255667469",
				spendTime:        180,
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

// 群聊里面抢红包的消息
func TestRedPacketGrabPushToGroup(t *testing.T) {
	type args struct {
		OperateID         string
		SendPacketUserID  string
		ClickPacketUserID string
		RedPacketID       string
		SendUserName      string
		ClickUserName     string
		GroupID           string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试给群推送抢红包消息",
			args: args{
				OperateID:         "123",
				SendPacketUserID:  "1914080869", //
				ClickPacketUserID: "1914080869", //
				RedPacketID:       "123",
				SendUserName:      "steven2",
				ClickUserName:     "steven2",
				GroupID:           "3483462779",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RedPacketGrabPushToGroup(tt.args.OperateID, tt.args.SendPacketUserID, tt.args.ClickPacketUserID, tt.args.RedPacketID, tt.args.SendUserName, tt.args.ClickUserName, tt.args.GroupID); (err != nil) != tt.wantErr {
				t.Errorf("RedPacketGrabPushToGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 测试给用户发送消息
func TestRedPacketGrabPushToUser(t *testing.T) {
	type args struct {
		OperateID         string
		SendMessageUserID string
		SendPacketUserID  string
		RedPacketID       string
		SendUserName      string
		ClickUserName     string
		ReceiveID         string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试给用户发送消息",
			args: args{
				OperateID:         "123",
				SendMessageUserID: "967320420",
				SendPacketUserID:  "1914080869",
				RedPacketID:       "123",
				SendUserName:      "steven2",
				ClickUserName:     "steven2",
				ReceiveID:         "1914080869",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RedPacketGrabPushToUser(tt.args.OperateID, tt.args.SendMessageUserID, tt.args.SendPacketUserID, tt.args.RedPacketID, tt.args.SendUserName, tt.args.ClickUserName, tt.args.ReceiveID); (err != nil) != tt.wantErr {
				t.Errorf("RedPacketGrabPushToUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
