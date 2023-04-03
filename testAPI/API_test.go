package testAPI

import (
	"Open_IM/pkg/base_info/notify"
	"encoding/json"
	"testing"
)

const BaseURL = "http://127.0.0.1:10002"

func TestPostAPI(t *testing.T) {
	type args struct {
		url       string
		construct func() []byte
	}
	tests := []struct {
		name       string
		args       args
		httpCode   int
		resultCode int
	}{
		{
			name: "提现回调接口:参数是不存在的:走成功逻辑",
			args: args{
				url: BaseURL + "/cloudWallet/charge_account_callback",
				construct: func() []byte {
					req := notify.ChargeNotifyReq{
						Version:         "1.0.0",
						TranCode:        "1001",
						MerOrderId:      "1234567890",
						MerId:           "1234567890",
						MerAttach:       "",
						Charset:         "UTF-8",
						SignType:        "MD5",
						ResultCode:      "",
						ErrorCode:       "",
						ErrorMsg:        "",
						OrderId:         "10086",
						TranAmount:      "",
						SubmitTime:      "",
						TranFinishTime:  "",
						BusinessType:    "",
						FeeAmount:       "",
						BankOrderId:     "",
						RealBankOrderId: "",
						DivideAcctDtl:   "",
						SignValue:       "",
					}
					content, err := json.Marshal(req)
					if err != nil {
						panic(err)
					}
					return content
				},
			},
			httpCode:   200,
			resultCode: 200,
		},
		{
			name: "转账失败回调：走失败回调逻辑",
			args: args{
				url: BaseURL + "/cloudWallet/charge_account_callback",
				construct: func() []byte {
					req := notify.ChargeNotifyReq{
						Version:    "1.0.0",
						TranCode:   "1001",
						MerOrderId: "1234567890",
						MerId:      "1234567890",
						MerAttach:  "",
						Charset:    "UTF-8",
						SignType:   "MD5",
						ResultCode: "",
						ErrorCode:  "4444",
						ErrorMsg:   "",
					}
					content, err := json.Marshal(req)
					if err != nil {
						panic(err)
					}
					return content
				},
			},
			httpCode:   200,
			resultCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpcode, errcode, err := PostAPI(tt.args.url, tt.args.construct)
			if err != nil {
				t.Errorf("PostAPI() error = %v", err)
				return
			}
			if httpcode != tt.httpCode {
				t.Errorf("PostAPI() httpcode = %v, want %v", httpcode, tt.httpCode)
			}
			if errcode != tt.resultCode {
				t.Errorf("PostAPI() errcode = %v, want %v", errcode, tt.resultCode)
			}
		})
	}
}
