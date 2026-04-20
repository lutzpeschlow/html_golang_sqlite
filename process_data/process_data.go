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
	// (1) options
	// (1.1) date
	fmt.Println(obj.Option1)
	t, err := parseGermanDate(obj.Option1)
	if err != nil {
		fmt.Println("ERROR: Date is not valid, please correct ...")
		return 1
	}
	fmt.Println("processing date: ", t)
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
	fmt.Println(obj.Scores[1])
	// return value
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
