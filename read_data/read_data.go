package read_data

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/lutzpeschlow/html_golang_sqlite/objects"
)

// ======================================================================================
//
// # ReadInput
//
// read values from form and store into data object
//
// ======================================================================================
// readInput liest Form-Daten in input, validiert und setzt Errors
func ReadInput(input *objects.InputData, r *http.Request, w http.ResponseWriter, tpl *template.Template) error {
	// options
	input.Option1 = r.FormValue("option1")
	input.Option2 = r.FormValue("option2")
	input.Option3 = r.FormValue("option3")
	input.Option4 = r.FormValue("option4")
	input.Option5 = r.FormValue("option5")
	// scores
	for i := 1; i <= 18; i++ {
		name := fmt.Sprintf("score%d", i)
		s := r.FormValue(name)
		if s == "" {
			input.Scores[i-1] = 0
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			input.Error = "invalid score: " + s
			tpl.Execute(w, input)
			return err // Stoppt Handler
		}
		if n < 0 || n > 99 {
			input.Error = "score must be 0–99"
			tpl.Execute(w, input)
			return fmt.Errorf("invalid range")
		}
		input.Scores[i-1] = n
	}
	return nil
}
