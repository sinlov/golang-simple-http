package simple_http

import (
	"reflect"
	"testing"
)

func TestPost_String(t *testing.T) {
	type args struct {
		urlStr     string
		headParams map[string](string)
		bodyParams map[string](string)
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
	}
	for _, tt := range tests {
		got, err := Post_String(tt.args.urlStr, tt.args.headParams, tt.args.bodyParams)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Post_String() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Post_String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestPost(t *testing.T) {
	type args struct {
		urlStr     string
		headParams map[string](string)
		bodyParams map[string](string)
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Post(tt.args.urlStr, tt.args.headParams, tt.args.bodyParams)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Post() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Post() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGet_String(t *testing.T) {
	type args struct {
		urlStr     string
		headParams map[string](string)
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Get_String(tt.args.urlStr, tt.args.headParams)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Get_String() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Get_String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGet(t *testing.T) {
	type args struct {
		urlStr     string
		headParams map[string](string)
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Get(tt.args.urlStr, tt.args.headParams)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Get() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Get() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGet_Sort(t *testing.T) {
	type args struct {
		urlStr     string
		headParams map[string](string)
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Get_Sort(tt.args.urlStr, tt.args.headParams)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Get_Sort() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Get_Sort() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
