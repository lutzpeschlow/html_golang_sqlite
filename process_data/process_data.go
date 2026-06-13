package process_data

import (
	"fmt"
	"strings"
	"time"

	"github.com/lutzpeschlow/html_golang_sqlite/objects"
)

// ======================================================================================
//
// # ProcessInputData
//
// the data from html template is processed and checked
// - are values valid
// - if there are incomplete values - update to valid values
// - no value input - update to default values
//
// with updated data the collection can be transferred and stored to database
//
// options:
// - date
// - location
// - type
// - any further
// - any further 2
//
// score:
// - 9 or 18 scores should be valid
//
// ======================================================================================
func ProcessInputData(obj *objects.InputData) int {
	fmt.Println("...")
	fmt.Println("data processing ...")

	// check on keyword mode
	keyword := strings.ToUpper(strings.TrimSpace(obj.Option1))
	keywords := []string{"CONTENT", "TEST", "HELP", "STATS"}
	// valid keyword - check in string slice
	isKeyword := false
	for _, kw := range keywords {
		if keyword == kw {
			isKeyword = true
			break
		}
	}
	// keyword valid - execute according function
	if isKeyword {
		switch keyword {
		case "CONTENT":
			fmt.Println("CONTENT Keyword erkannt")
			return ProcessContent(obj)
		case "TEST":
			fmt.Println("TEST Keyword erkannt")
			return ProcessTest(obj)
		case "HELP":
			fmt.Println("HELP Keyword erkannt")
			return ProcessHelp(obj)
		case "STATS":
			fmt.Println("STATS Keyword erkannt")
			return ProcessStats(obj)
		}
	}

	// (1) options
	// (1.1) date or option
	//       possible options:
	//           CONTENT
	fmt.Println("date parsing: ", obj.Option1)
	t, err := parseGermanDate(obj.Option1)
	if err != nil {
		fmt.Println("ERROR: Date is not valid, please correct ...")
		obj.Error = "Date is not valid"
		return 1
	}
	fmt.Println("processing date: ", t)
	obj.Date = t
	// (1.2) location
	obj.Option2 = strings.ToUpper(strings.TrimSpace(obj.Option2))
	// (1.3) Type
	obj.Option3 = strings.ToUpper(strings.TrimSpace(obj.Option3))
	if obj.Option3 != "PGA" && obj.Option3 != "CH" {
		obj.Option3 = "CH"
	}
	// (1.4) (1.5) further options, currently not processed
	obj.Option4 = strings.ToUpper(strings.TrimSpace(obj.Option4))
	obj.Option5 = strings.ToUpper(strings.TrimSpace(obj.Option5))
	// (2) Scores
	num_holes := countNonZero(obj.Scores)
	fmt.Println("non-zero holes: ", num_holes)
	if num_holes != 9 && num_holes != 18 {
		fmt.Println("ERROR: number of holes should be 9 or 18 ...")
		obj.Error = "Number of holes not valid"
		return 1
	}
	fmt.Println("finalized processing ...")
	// return value
	return 0

}

func ProcessHelp(obj *objects.InputData) int {
	obj.Result = "Verfügbare Keywords: CONTENT, TEST, HELP, STATS"
	obj.Error = ""
	return 0
}

func ProcessStats(obj *objects.InputData) int {
	obj.Result = "Statistik wird berechnet..."
	// Hier Statistik-Logik
	return 0
}
func ProcessContent(obj *objects.InputData) int {
	obj.Result = "Verfügbare Keywords: CONTENT, TEST, HELP, STATS"
	obj.Error = ""
	return 0
}

func ProcessTest(obj *objects.InputData) int {
	obj.Result = "Statistik wird berechnet..."
	// Hier Statistik-Logik
	return 0
}

// ======================================================================================
//
// parseGermanDate
//
// input format as german format like 10.4.2026
// will be converted into time.Time for further processing
//
// ======================================================================================
func parseGermanDate(s string) (time.Time, error) {
	s = strings.TrimSpace(s)

	layouts := []string{
		"2.1.2006",
		"02.01.2006",
	}

	var lastErr error
	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
		lastErr = err
	}
	return time.Time{}, lastErr
}

// ======================================================================================
//
// countNonZero
//
// the array of
//
// ======================================================================================
func countNonZero(scores [18]int) int {
	count := 0
	for _, v := range scores {
		if v != 0 {
			count++
		}
	}
	return count
}
