package system

import (
	"fmt"
	"os"
	"runtime"
)

type System struct {
	architecture    string
	operatingSystem string
	drives          []*Drive
}

func NewSystem() *System {
	newSystem := System{runtime.GOARCH, runtime.GOOS, nil}
	return &newSystem
}

func (s *System) Print() {
	fmt.Printf("Architecture:      %s\n", s.architecture)
	fmt.Printf("OperatingSystem:   %s\n", s.operatingSystem)
	fmt.Println("Drives:")
	for _, drive := range s.drives {
		drive.ScanForFoldersAndFiles()
		drive.Print()
	}
}

func (s *System) GetDrives() {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		driveRoot := string(drive) + ":\\"
		f, err := os.Open(driveRoot)
		if err == nil {
			s.drives = append(s.drives, NewDrive(driveRoot))
			f.Close()
		}
	}
	return
}

func (s *System) ScanForFoldersAndFiles() {
	for _, drive := range s.drives {
		drive.ScanForFoldersAndFiles()
		drive.Print()
	}
}
