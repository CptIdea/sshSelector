package ssh

import (
	"embed"
	"reflect"
	"testing"
)

//go:embed test_file
var _ embed.FS

func TestNewManagerFromFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    Manager
		wantErr bool
	}{
		{"file not found",args{file: "sndljzxcljawskcnnx"},nil,true},
		{"create",args{file: "test_file"},&fromFileManager{[]string{"root@example.com"}},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewManagerFromFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewManagerFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManagerFromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fromFileManager_GetList(t *testing.T) {
	type fields struct {
		list []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{"get list",fields{list: []string{"test1","test2"}},[]string{"test1","test2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fromFileManager{
				list: tt.fields.list,
			}
			if got := f.GetList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}
