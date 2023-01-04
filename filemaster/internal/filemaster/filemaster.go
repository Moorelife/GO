package filemaster

import (
	"runtime"

	"github.com/Moorelife/GO/filemaster/pkg/system"
)

type FileMaster struct {
	system.System
}

func NewFileMaster() *FileMaster {
	instance := FileMaster{
		system.System(runtime.GOARCH, runtime.GOOS, nil),
	}
	return &instance
}

var DriveLetters = []string{}

func (f *FileMaster) RecognizeDrives() {

}
