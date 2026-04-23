package objects

import (
	"fmt"
	"time"
)

type InputData struct {
	Date    time.Time
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
	Score1    int
	Score2    int
	Score3    int
	Score4    int
	Score5    int
	Score6    int
	Score7    int
	Score8    int
	Score9    int
	Score10   int
	Score11   int
	Score12   int
	Score13   int
	Score14   int
	Score15   int
	Score16   int
	Score17   int
	Score18   int
	CreatedAt time.Time
}

func GetDebugPrintoutInput(obj *InputData) {
	fmt.Println("debug printout of InputData:")
	fmt.Printf("%+v\n", obj)
}

func GetDebugPrintoutSql(obj *SqlData) {
	fmt.Println("debug printout of SqlData:")
	fmt.Printf("%+v\n", obj)
}
