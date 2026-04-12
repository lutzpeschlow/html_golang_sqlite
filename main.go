package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

type Data struct {
	Num1   int
	Num2   int
	Result string
	Error  string
}

func main() {
	// Template laden
	tpl := template.Must(template.ParseFiles("templates/index.html"))

	// Routes direkt registrieren (KEINE HandlerFunc-Wrapper)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, Data{})
	})

	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

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

		result := processIntegers(num1, num2)
		tpl.Execute(w, Data{Num1: num1, Num2: num2, Result: result})
	})

	fmt.Println("🚀 Server: http://localhost:8080")
	fmt.Println("⏳ Öffne Browser...")

	// Browser nach 3s (Linux Mint Vivaldi)
	go func() {
		time.Sleep(3 * time.Second)
		exec.Command("vivaldi-stable", "http://localhost:8080").Start()
	}()

	// BLOCKIEREND - Server läuft
	http.ListenAndServe(":8080", nil)
}

func processIntegers(a, b int) string {
	return fmt.Sprintf(
		"Ergebnisse:\n• Summe: %d + %d = %d\n• Produkt: %d × %d = %d\n• Differenz: %d - %d = %d\n• Quotient: %d ÷ %d = %d",
		a, b, a+b, a, b, a*b, a, b, a-b, a, b, a/b,
	)
}
