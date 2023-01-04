// FileMaster is a program that scans all local drives, and finds all unique versions
// of all files available on these drives. In time it will be extended to cleverly
// reorganize those folders the user wants.
package main

import (
	"fmt"

	"github.com/Moorelife/GO/filemaster/internal/filemaster"
)

func main() {
	fmt.Println("FileMaster version 0.1")

	Master := filemaster.NewFileMaster()
	Master.RecognizeDrives()
	// Master.ScanForFoldersAndFiles()
	Master.Print()

}
