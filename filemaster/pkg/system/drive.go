package system

import (
	"fmt"
	"os"
)

type Drive struct {
	driveRoot string
	isPresent bool
}

func NewDrive(driveRoot string) *Drive {
	newDrive := Drive{driveRoot, false}
	f, err := os.Open(driveRoot)
	if err == nil {
		f.Close()
		newDrive.isPresent = true
	}
	return &newDrive
}

func (d *Drive) Print() {
	fmt.Printf("Drive %#v ,present: %#v", d.driveRoot, d.isPresent)
}
