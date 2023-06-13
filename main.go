package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Booting up Hamster-Land backend server.")

	fs := http.FileServer(http.Dir("./static/"))
	h := http.NewServeMux()
	h.Handle("/static/", http.StripPrefix("/static/", fs))
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving 'tmpls/index.html'.")
		b, err := os.ReadFile("tmpls/index.html")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Fprint(w, string(b))
	})
	h.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving 'tmpls/about.html'.")
		b, err := os.ReadFile("tmpls/about.html")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Fprint(w, string(b))
	})
	h.HandleFunc("/news/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving 'tmpls/news.html'.")
		b, err := os.ReadFile("tmpls/news.html")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Fprint(w, string(b))
	})
	h.HandleFunc("/me/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving 'tmpls/me.html'.")
		b, err := os.ReadFile("tmpls/me.html")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Fprint(w, string(b))
	})

	err := http.ListenAndServe(":8080", h)
	if err != nil {
		fmt.Print(err)
	}
}
