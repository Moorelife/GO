package filemaster

import (
	"github.com/Moorelife/GO/filemaster/pkg/system"
)

type FileMaster struct {
	system *system.System
}

func NewFileMaster() *FileMaster {
	instance := FileMaster{
		system: system.NewSystem(),
	}
	return &instance
}

var DriveLetters = []string{}

func (f *FileMaster) RecognizeDrives() {
	f.system.GetDrives()
}

func (f *FileMaster) ScanForFoldersAndFiles() {
	f.system.ScanForFoldersAndFiles()
}

func (f *FileMaster) Print() {
	f.system.Print()
}
