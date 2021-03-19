package apple_login

import (
	"testing"
)

func TestVerifyToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				token: "eyJraWQiOiJZdXlYb1kiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiY29tLnllc2ZlZWxpbmcuZml0bmVzc1oiLCJleHAiOjE2MTYyMjM2NjQsImlhdCI6MTYxNjEzNzI2NCwic3ViIjoiMDAwNDAyLjY5ZWMyOGU4M2NiMTQ3NDdhNzE1N2U4M2FlZDAzYmJmLjA2NTQiLCJjX2hhc2giOiI2dE5jb2pLdUFyTU90ZlpQWWtQVjJBIiwiZW1haWwiOiJtNTdreGJiZzk5QHByaXZhdGVyZWxheS5hcHBsZWlkLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjoidHJ1ZSIsImlzX3ByaXZhdGVfZW1haWwiOiJ0cnVlIiwiYXV0aF90aW1lIjoxNjE2MTM3MjY0LCJub25jZV9zdXBwb3J0ZWQiOnRydWV9.IOnYN0WnfiZHDSF2gwFbnKQqWvqxnNxvlsEsUkH88MdBNoxbPQFISq8rLIE4LZIMehOiYvSpdkLZIVTVa2kXhFob38BGXw9JCVxFfjgBXDZou7Dma9HZ_WGybJvTLp2Iz50HKVG_N3EsyjBNr6TUp9DBZrwYeeXalzWKH-_Taj11LvRWYYqoV6CCQ8D-UYI_MqNOlVd5WVtZhoQQKVB0HMK0iZFK_gc1jmUD6IqCtIAs5GRwrPQ_lw6--Ja_dx7rofyH3c4OjhTwZz4Hm6NxOIoMAaAsHzR0aEb6A1qHt944DOTCotrS5Xi13VU2yOFUb7IilhoC0beXROsjLToYlg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyToken(tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVerifyToken2(t *testing.T) {
	type args struct {
		payload string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				payload: "eyJraWQiOiJZdXlYb1kiLCJhbGciOiJSUzI1NiJ9.eyJpc3MiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiYXVkIjoiY29tLnllc2ZlZWxpbmcuZml0bmVzc1oiLCJleHAiOjE2MTYyMjM2NjQsImlhdCI6MTYxNjEzNzI2NCwic3ViIjoiMDAwNDAyLjY5ZWMyOGU4M2NiMTQ3NDdhNzE1N2U4M2FlZDAzYmJmLjA2NTQiLCJjX2hhc2giOiI2dE5jb2pLdUFyTU90ZlpQWWtQVjJBIiwiZW1haWwiOiJtNTdreGJiZzk5QHByaXZhdGVyZWxheS5hcHBsZWlkLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjoidHJ1ZSIsImlzX3ByaXZhdGVfZW1haWwiOiJ0cnVlIiwiYXV0aF90aW1lIjoxNjE2MTM3MjY0LCJub25jZV9zdXBwb3J0ZWQiOnRydWV9.IOnYN0WnfiZHDSF2gwFbnKQqWvqxnNxvlsEsUkH88MdBNoxbPQFISq8rLIE4LZIMehOiYvSpdkLZIVTVa2kXhFob38BGXw9JCVxFfjgBXDZou7Dma9HZ_WGybJvTLp2Iz50HKVG_N3EsyjBNr6TUp9DBZrwYeeXalzWKH-_Taj11LvRWYYqoV6CCQ8D-UYI_MqNOlVd5WVtZhoQQKVB0HMK0iZFK_gc1jmUD6IqCtIAs5GRwrPQ_lw6--Ja_dx7rofyH3c4OjhTwZz4Hm6NxOIoMAaAsHzR0aEb6A1qHt944DOTCotrS5Xi13VU2yOFUb7IilhoC0beXROsjLToYlg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyToken2(tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("Verify2() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
