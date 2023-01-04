package system

import (
	"reflect"
	"testing"
)

func TestNewDrive(t *testing.T) {
	type args struct {
		driveRoot string
	}
	tests := []struct {
		name string
		args args
		want *Drive
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDrive(tt.args.driveRoot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDrive() = %v, want %v", got, tt.want)
			}
		})
	}
}
