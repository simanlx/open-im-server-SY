package ncount

import (
	"reflect"
	"testing"
)

func TestEncrpt(t *testing.T) {
	type args struct {
		message         []byte
		publicKeyString string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test encrpt",
			args: args{
				message:         []byte("test"),
				publicKeyString: PUBLIC_KEY,
			},
			want:    []byte("test"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrpt(tt.args.message, tt.args.publicKeyString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrpt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrpt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
