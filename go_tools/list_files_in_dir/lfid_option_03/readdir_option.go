package main

import (
	"fmt"
	"os"
)

var targetPath string = "/Users/jasoncameron/00_Drive/Core/00_Managers/03_RPG_Management/RPGMS/01_DnD/02_DnD_Assets/Items/Jewelry"

func main() {
	f, err := os.Open(targetPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := f.ReadDir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		fmt.Println(v.Name(), v.IsDir())
	}

}
