package ncount

import (
	"reflect"
	"testing"
)

func TestNewCounter(t *testing.T) {
	tests := []struct {
		name string
		want NCounter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCounter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_BindCard(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *BindCardReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *BindCardResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.BindCard(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("BindCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BindCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_BindCardConfirm(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *BindCardConfirmReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *BindCardConfirmResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.BindCardConfirm(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("BindCardConfirm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BindCardConfirm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_CheckUserAccountDetail(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *CheckUserAccountDetailReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CheckUserAccountDetailResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.CheckUserAccountDetail(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUserAccountDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckUserAccountDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_CheckUserAccountInfo(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *CheckUserAccountReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CheckUserAccountResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.CheckUserAccountInfo(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUserAccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckUserAccountInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_CheckUserAccountTrans(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *CheckUserAccountTransReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CheckUserAccountTransResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.CheckUserAccountTrans(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUserAccountTrans() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckUserAccountTrans() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_NewAccount(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *NewAccountReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *NewAccountResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.NewAccount(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_QuickPayConfirm(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *QuickPayConfirmReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *QuickPayConfirmResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.QuickPayConfirm(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuickPayConfirm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickPayConfirm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_QuickPayOrder(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *QuickPayOrderReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *QuickPayOrderResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.QuickPayOrder(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuickPayOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickPayOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_Refund(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *RefundReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RefundResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.Refund(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Refund() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Refund() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_Transfer(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *TransferReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TransferResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.Transfer(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transfer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_UnbindCard(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *UnBindCardReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UnBindCardResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.UnbindCard(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnbindCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnbindCard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counter_Withdraw(t *testing.T) {
	type fields struct {
		notifyQuickPayConfirmURL string
		notifyRefundURL          string
		notifyWithdrawURL        string
	}
	type args struct {
		req *WithdrawReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *WithdrawResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &counter{
				notifyQuickPayConfirmURL: tt.fields.notifyQuickPayConfirmURL,
				notifyRefundURL:          tt.fields.notifyRefundURL,
				notifyWithdrawURL:        tt.fields.notifyWithdrawURL,
			}
			got, err := c.Withdraw(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Withdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Withdraw() got = %v, want %v", got, tt.want)
			}
		})
	}
}
