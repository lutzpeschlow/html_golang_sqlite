package objects

import (
	"fmt"
	"reflect"
	"time"
)

// Control stores the control file content.
type Control struct {
	Enable EnableConfig `json:"enable"`
	DbPath string       `json:"dbpath"`
}

// EnableConfig stores the main enable switches.
type EnableConfig struct {
	WRITE       bool        `json:"WRITE"`
	READ        bool        `json:"READ"`
	READOptions ReadOptions `json:"READ_OPTIONS"`
}

// ReadOptions stores sub-options for READ.
type ReadOptions struct {
	COMMON     bool `json:"COMMON"`
	NumEntries bool `json:"NUM_ENTRIES"`
}

// store input from html
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

// store prepared sql data
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

// content of sql
type SqlContent struct {
	DbPath string
	Count  int64
	Dates  []string
}

// stats of sql

// ======================================================================================
//
//	DebugPrintout
//
// debug printout of object with all details
//
// ======================================================================================
func DebugPrintout(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	fmt.Printf("%s: \n", t.Name())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if !field.IsExported() || (value.IsZero() && value.Kind() != reflect.String) {
			continue
		}
		fmt.Printf("  %s: %v\n", field.Name, value.Interface())
	}
}
