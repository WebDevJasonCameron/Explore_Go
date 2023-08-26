package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var path string = "/Users/jasoncameron/00_Drive/Core/00_Managers/03_RPG_Management/RPGMS/01_DnD/02_DnD_Assets/Items/Armor"

func main() {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}
