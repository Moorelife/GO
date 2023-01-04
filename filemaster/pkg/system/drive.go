package system

import (
	"fmt"
	"os"
)

type Drive struct {
	driveRoot string
}

func NewDrive(driveRoot string) *Drive {
	newDrive := Drive{driveRoot}
	f, err := os.Open(driveRoot)
	if err == nil {
		f.Close()
	}
	return &newDrive
}

func (d *Drive) Print() {
	fmt.Printf("%v\n", d.driveRoot)
}
