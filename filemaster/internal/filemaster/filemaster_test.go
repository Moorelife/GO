package filemaster

import (
	"runtime"
	"testing"
)

func TestNewFileMaster(t *testing.T) {
	fm := NewFileMaster()
	if fm == nil {
		t.Error("NewFileMaster returned nil")
	}
	if fm.System.Architecture != runtime.GOARCH {
		t.Errorf("NewFileMaster should have Architecture %s\n", runtime.GOARCH)
	}
	if fm.System.OperatingSystem != runtime.GOOS {
		t.Errorf("NewFileMaster should have Operating System %s\n", runtime.GOOS)
	}
	if len(fm.System.Drives) > 0 {
		t.Errorf("NewFileMaster should not have any Drives yet, but has %v", fm.System.Drives)
	}
}
