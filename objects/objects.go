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

func GetDebugPrintout(obj *InputData) {
	fmt.Println("debug printout ...")
	fmt.Println("  ", obj.Option1)
	fmt.Println("  ", obj.Option2)
	fmt.Println("  ", obj.Option3)
	fmt.Println("  ", obj.Option4)
	fmt.Println("  ", obj.Option5)
	fmt.Println("  ", obj.Scores)

}
