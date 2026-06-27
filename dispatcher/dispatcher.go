package dispatcher

import (
	"fmt"

	"github.com/lutzpeschlow/html_golang_sqlite/ctrl"
	"github.com/lutzpeschlow/html_golang_sqlite/html_input"
	"github.com/lutzpeschlow/html_golang_sqlite/objects"
	"github.com/lutzpeschlow/html_golang_sqlite/sql_data"
)

func ExecuteAction(c *objects.Control) error {
	var enabled string
	fmt.Println(" execute action ....")

	enabled = ctrl.GetEnabled(c)
	fmt.Println(enabled)

	switch enabled {
	case "WRITE":
		fmt.Println("write is on, open html file for input data into database")
		status := html_input.HtmlInput(c.DbPath)
		if status != nil {
			fmt.Println("error")
		}
	case "READ":
		fmt.Println("read is on, use get content function to read database")
		switch {
		case c.Enable.READOptions.COMMON:
			fmt.Println("READ option COMMON is on")
			if err := sql_data.GetContentCommon(c.DbPath); err != nil {
				fmt.Println("error:", err)
				return err
			}

		case c.Enable.READOptions.NumEntries:
			fmt.Println("READ option NUM_ENTRIES is on")
			if err := sql_data.GetNumEntries(c.DbPath); err != nil {
				fmt.Println("error:", err)
				return err
			}

		default:
			fmt.Println("no READ option enabled")
		}
	default:
		fmt.Println("no action enabled")
	}

	return nil
}
