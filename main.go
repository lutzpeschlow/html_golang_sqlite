package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Data struct {
	Num1   int
	Num2   int
	Result string
	Error  string
}

func main() {
	// Template laden (separate index.html)
	tpl := template.Must(template.ParseFiles("templates/index.html"))

	// Hauptseite (GET)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, Data{})
	})

	// POST-Handler
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		// Formular parsen
		num1Str := r.FormValue("num1")
		num2Str := r.FormValue("num2")

		num1, err := strconv.Atoi(num1Str)
		if err != nil {
			tpl.Execute(w, Data{Error: "Erste Zahl ungültig!"})
			return
		}

		num2, err := strconv.Atoi(num2Str)
		if err != nil {
			tpl.Execute(w, Data{Error: "Zweite Zahl ungültig!"})
			return
		}

		// Go-Funktion
		result := processIntegers(num1, num2)

		// Zurück mit Ergebnis
		tpl.Execute(w, Data{Num1: num1, Num2: num2, Result: result})
	})

	fmt.Println("🚀 Server: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Deine Funktion
func processIntegers(a, b int) string {
	return fmt.Sprintf(
		"Ergebnisse:\n"+
			"• Summe: %d + %d = %d\n"+
			"• Produkt: %d × %d = %d\n"+
			"• Differenz: %d - %d = %d\n"+
			"• Quotient: %d ÷ %d = %d",
		a, b, a+b,
		a, b, a*b,
		a, b, a-b,
		a, b, a/b,
	)
}
