package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

var targetPath string = "/Users/jasoncameron/00_Drive/Core/00_Managers/03_RPG_Management/RPGMS/01_DnD/02_DnD_Assets/Items"

func main() {
	err := filepath.Walk(targetPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
