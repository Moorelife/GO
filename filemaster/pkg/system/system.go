package system

import (
	"os"
)

type System struct {
	architecture    string
	operatingSystem string
	drives          []*Drive
}

func NewSystem(architecture, operatingSystem string, drives []*Drive) *System {
	newSystem := System{architecture, operatingSystem, drives}
	return &newSystem
}

func (s *System) Print() {

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
