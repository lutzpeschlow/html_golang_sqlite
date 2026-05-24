package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/lutzpeschlow/html_golang_sqlite/ctrl"
	"github.com/lutzpeschlow/html_golang_sqlite/dispatcher"
	"github.com/lutzpeschlow/html_golang_sqlite/objects"
)

// ======================================================================================
// main
//
// main function
//
// - call ctrl package to read json file and fill ctrl object
//
// ======================================================================================
func main() {
	var control objects.Control
	var enabled string
	// (1) read control json file
	if err := ctrl.ReadControlJsonFile("control.json", &control); err != nil {
		fmt.Println("control error:", err)
		return
	}
	// (2) dispatcher spreads further actions
	// - WRITE
	// - READ
	enabled = ctrl.GetEnabled(&control)
	fmt.Println(enabled)
	dispatcher.ExecuteAction(&control)

}
