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

type Data struct {
	Num1   int
	Num2   int
	Result string
	Error  string
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = tpl.Execute(w, Data{})
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
		// convert into integer
		num1Str := r.FormValue("num1")
		num2Str := r.FormValue("num2")
		num1, err := strconv.Atoi(num1Str)
		if err != nil {
			_ = tpl.Execute(w, Data{Error: "invalid number"})
			return
		}
		num2, err := strconv.Atoi(num2Str)
		if err != nil {
			_ = tpl.Execute(w, Data{Error: "invalid number"})
			return
		}
		// processing numbers in separate function
		result := processIntegers(num1, num2)
		// fill template with numbers
		_ = tpl.Execute(w, Data{Num1: num1, Num2: num2, Result: result})
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
// processIntegers
//
// ======================================================================================
func processIntegers(a, b int) string {
	fmt.Println("data processing ...")
	fmt.Println(a, b)
	return "result"
}
