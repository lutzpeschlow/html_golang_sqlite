package html_input

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/lutzpeschlow/html_golang_sqlite/objects"
	"github.com/lutzpeschlow/html_golang_sqlite/process_data"
	"github.com/lutzpeschlow/html_golang_sqlite/sql_data"
)

// ======================================================================================
//
//	HtmlInput
//
// - (1) load and checks template
// - (2) register route /, if the browser calls URL, handler will be executed
// - (3) rendering template and write output into http response
// - (4) register second route for formular posting, reacts in case of /calc
// - (5) read data
// - (6) process data
// - (7) [output template again with changed values]
// - (8) starts webserver, block main thread
//
// ======================================================================================
func HtmlInput() error {
	// (1) load html template from index.html
	// read the file with parsefiles and create a template object
	// must valid template
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	// register a route
	// with / start function, w is the response, r is the request from browser
	// anonymous function, will defined directly
	// tpl.execute is rendering template and it will send to browser
	// data object as empty object
	// inp_obj := objects.InputData{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tpl.Execute(w, objects.InputData{})
	})
	// register second route, responsible for final content
	// send formular per POST
	// if no POST, e.g. GET no response and return
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		// read input lines
		var input objects.InputData
		var sql objects.SqlData
		if err := ReadInput(&input, r, w, tpl); err != nil {
			return
		}
		// processing numbers in separate function
		status := process_data.ProcessInputData(&input)
		fmt.Println("in main after processing ...")
		if status > 0 {
			fmt.Println("ERROR: no valid input ...")
			input.Error = "Invalid input data"
			_ = tpl.Execute(w, input)
			return
		}
		objects.DebugPrintout(input)
		// transfer to sql object
		status = sql_data.PrepareSQLData(&input, &sql)
		objects.DebugPrintout(sql)
		// fill template with numbers
		_ = tpl.Execute(w, input)

		db, err := sql_data.CreateDB("./data.sqlite")
		if err != nil {
			fmt.Println("DB error:", err)
			return
		}
		defer db.Close()

		if err := sql_data.InitSchema(db); err != nil {
			fmt.Println("schema error:", err)
			return
		}

		id, err := sql_data.InsertSQLData(db, sql)
		if err != nil {
			fmt.Println("insert error:", err)
			return
		}
		fmt.Println("Inserted ID:", id)
	})
	// adress of server
	fmt.Println("Server: http://localhost:8080")
	fmt.Println("Browser...")
	// start go routine
	// parallel running without blocking main
	// wait for three seconds, time for webserver before opening browser with adress
	go func() {
		time.Sleep(3 * time.Second)

		if err := openBrowser("http://localhost:8080"); err != nil {
			fmt.Println("  browser error: ", err)
		}
	}()
	// webserver start at port 8080
	_ = http.ListenAndServe(":8080", nil)

	return nil

}

// ======================================================================================
//
// openBrowser
//
// depending on used operating system the start of browser is different
//
// ======================================================================================
func openBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("vivaldi-stable", "--new-tab", url).Start()
	case "windows":
		return exec.Command("cmd", "/c", "start", "msedge", url).Start()
	default:
		return exec.Command("xdg-open", url).Start()
	}
}

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
