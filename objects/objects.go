package objects

import "fmt"

type InputData struct {
	Option1 string
	Option2 string
	Option3 string
	Option4 string
	Option5 string
	Scores  [18]int
	Result  string
	Error   string
}

func GetDebugPrintout(o *InputData) {
	fmt.Println("debug printout of InputData:")
	fmt.Println(" ", o.Option1, ",", o.Option2, ",", o.Option3, ",", o.Option4, ",", o.Option5)
	fmt.Println(" ", o.Scores)

}
