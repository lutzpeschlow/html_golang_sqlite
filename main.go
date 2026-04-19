package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

type InputData struct {
	Option1, Option2, Option3, Option4, Option5 string
	Scores                                      [18]int
	Result                                      string
	Error                                       string
}

// ======================================================================================
// main
//
// main function
// - load template
// - register two routines
// - start server
//
// ======================================================================================
func main() {
	// load html template from index.html
	// read the file with parsefiles and create a template object
	// must valid template
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	// register a route
	// with / start function, w is the response, r is the request from browser
	// anonymous function, will defined directly
	// tpl.execute is rendering template and it will send to browser
	// data object as empty object
	inp_obj := InputData{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tpl.Execute(w, inp_obj)
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
		var input InputData
		if err := readInput(&input, r, w, tpl); err != nil {
			return
		}
		// processing numbers in separate function
		processInput(&input)
		// fill template with numbers
		_ = tpl.Execute(w, input)
	})
	// adress of server
	fmt.Println("Server: http://localhost:8080")
	fmt.Println("Browser...")
	// start go routine
	// parallel running without blocking main
	// wait for three seconds, time for webserver before opening browser with adress
	go func() {
		time.Sleep(3 * time.Second)
		openBrowser("http://localhost:8080")
	}()
	// webserver start at port 8080
	_ = http.ListenAndServe(":8080", nil)
}

// ======================================================================================
//
// openBrowser
//
// ======================================================================================
func openBrowser(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("vivaldi-stable", url).Start()
	case "windows":
		return exec.Command("cmd", "/c", "start", "msedge", url).Start()
	default:
		return exec.Command("xdg-open", url).Start()
	}
}

// ======================================================================================
//
// readInput
//
// ======================================================================================
// readInput liest Form-Daten in input, validiert und setzt Errors
func readInput(input *InputData, r *http.Request, w http.ResponseWriter, tpl *template.Template) error {
	// 5 Optionen als Strings
	input.Option1 = r.FormValue("option1")
	input.Option2 = r.FormValue("option2")
	input.Option3 = r.FormValue("option3")
	input.Option4 = r.FormValue("option4")
	input.Option5 = r.FormValue("option5")

	// 18 Scores validieren & parsen
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
	return nil // Erfolg
}

// ======================================================================================
//
// processIntegers
//
// ======================================================================================
func processInput(input *InputData) {
	fmt.Println("data processing ...")
	fmt.Println(input.Option1)
	fmt.Println(input.Scores[1])

}
