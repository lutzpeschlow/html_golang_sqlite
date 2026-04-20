package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/lutzpeschlow/html_golang_sqlite/objects"
	"github.com/lutzpeschlow/html_golang_sqlite/process_data"
	"github.com/lutzpeschlow/html_golang_sqlite/read_data"
)

// ======================================================================================
// main
//
// main function
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
func main() {
	// (1) load html template from index.html
	// read the file with parsefiles and create a template object
	// must valid template
	tpl := template.Must(template.ParseFiles("templates/index.html"))
	// register a route
	// with / start function, w is the response, r is the request from browser
	// anonymous function, will defined directly
	// tpl.execute is rendering template and it will send to browser
	// data object as empty object
	inp_obj := objects.InputData{}
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
		var input objects.InputData
		if err := read_data.ReadInput(&input, r, w, tpl); err != nil {
			return
		}
		// processing numbers in separate function
		process_data.ProcessInputData(&input)
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
// depending on used operating system the start of browser is different
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
