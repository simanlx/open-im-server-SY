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
