package fileops

import (
	"os"
	"reflect"
	"testing"
)

func TestFileManager_Read(t *testing.T) {
	f, err := NewFileManager("abc.txt", 10)
	if err != nil {
		t.Errorf("Error in creating file manager: %v", err)
	}

	type fields struct {
		file     *os.File
		pageSize int
		numPages int
	}
	type args struct {
		pageNum int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Reading from file",
			fields: fields{
				file:     f.file,
				pageSize: f.pageSize,
				numPages: f.numPages,
			},
			args: args{
				pageNum: 0,
			},
			want:    []byte("Hello!!   "), // since pageSize is 10, three additional spaces are needed to satisfy the equality check.
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.Read(tt.args.pageNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileManager.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileManager.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileManager_Write(t *testing.T) {
	f, err := NewFileManager("abc.txt", 10)
	if err != nil {
		t.Errorf("Error in creating file manager: %v", err)
	}

	type fields struct {
		file     *os.File
		pageSize int
		numPages int
	}
	type args struct {
		data    []byte
		pageNum int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Adding a random string",
			fields: fields{
				file:     f.file,
				pageSize: f.pageSize,
				numPages: f.numPages,
			},
			args: args{
				data:    []byte("Hello!!"),
				pageNum: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := f.Write(tt.args.data, tt.args.pageNum); (err != nil) != tt.wantErr {
				t.Errorf("FileManager.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
