package main

import (
	"fmt"

	"github.com/fabianoflorentino/solid/dip"
)

func main() {

	// Database
	readDatabase := dip.Database{}
	serviceDatabase := dip.DataService{ReadData: readDatabase}
	fmt.Println(serviceDatabase.GetData())

	// File
	readFile := dip.File{}
	serviceFile := dip.DataService{ReadData: readFile}
	fmt.Println(serviceFile.GetData())
}
