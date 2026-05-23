package dispatcher

import (
	"fmt"

	"github.com/lutzpeschlow/html_golang_sqlite/objects"
)

func ExecuteAction(ctrl *objects.Control) error {
	fmt.Println(" execute action ....")

	// 	switch ctrl.Enable {
	// 	case "WRITE":
	// 		fmt.Println("write is on, open html file for input data into database")
	// 	case "READ":
	// 		fmt.Println("read is on, use get content function to read database")
	// 	}
	return nil
}
