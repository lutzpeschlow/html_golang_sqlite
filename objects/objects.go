package objects

import (
	"fmt"
	"time"
)

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

type SqlData struct {
	ID        int64
	Date      time.Time
	Option2   string
	Option3   string
	Option4   string
	Option5   string
	Scores    [18]int
	CreatedAt time.Time
}

func GetDebugPrintoutInput(o *InputData) {
	fmt.Println("debug printout of InputData:")
	fmt.Println(" ", o.Option1, ",", o.Option2, ",", o.Option3, ",", o.Option4, ",", o.Option5)
	fmt.Println(" ", o.Scores)
}

func GetDebugPrintoutSql(s *SqlData) {
	fmt.Println("debug printout of SqlData:")
	fmt.Println(s)
}
