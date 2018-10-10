package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println("---- File ----")
		fmt.Println("Name:", info.Name())       // base name of the file
		fmt.Println("Size:", info.Size())       // length in bytes for regular files; system-dependent for others
		fmt.Println("Mode:", info.Mode())       // file mode bits
		fmt.Println("ModTime:", info.ModTime()) // modification time
		fmt.Println("IsDir:", info.IsDir())     // abbreviation for Model().IsDir()
		fmt.Println("Sys:", info.Sys())         // underlying data source (can return nil)
		fmt.Println("--------------")
		return nil
	})
}
