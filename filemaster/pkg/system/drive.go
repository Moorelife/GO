package system

import (
	"fmt"
	"os"
	"path/filepath"
)

type Drive struct {
	driveRoot string
	folders   []string
	files     []string
}

func NewDrive(driveRoot string) *Drive {
	newDrive := Drive{driveRoot, nil, nil}
	f, err := os.Open(driveRoot)
	if err == nil {
		f.Close()
	}
	return &newDrive
}

func (d *Drive) ScanForFoldersAndFiles() error {
	err := filepath.Walk(d.driveRoot, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			d.files = append(d.files, path)
		} else {
			d.folders = append(d.folders, path)
		}
		return nil
	})
	return err
}

func (d *Drive) Print() {
	fmt.Printf("%v folders: %v  files: %v\n", d.driveRoot, len(d.folders), len(d.files))
}
