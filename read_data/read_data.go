package read_data

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/lutzpeschlow/html_golang_sqlite/objects"
	"github.com/lutzpeschlow/html_golang_sqlite/sql_data"
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

func UpdateContentHTML(w http.ResponseWriter, tpl *template.Template, dbPath string) error {
	// Datenbank öffnen
	db, err := sql_data.CreateDB(dbPath)
	if err != nil {
		return fmt.Errorf("could not open database: %w", err)
	}
	defer db.Close()

	count, err := sql_data.GetRoundCount(db)
	if err != nil {
		return fmt.Errorf("could not get count: %w", err)
	}
	dates, err := sql_data.GetAllDates(db)
	if err != nil {
		return fmt.Errorf("could not get dates: %w", err)
	}
	fmt.Println(dates)

	content := objects.SqlContent{
		Count: count,
		Dates: dates,
	}

	// Template mit den Daten rendern
	return tpl.ExecuteTemplate(w, "content.html", content)
}
