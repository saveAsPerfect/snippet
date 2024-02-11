package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var like int = 0

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		return
	}

	t.ExecuteTemplate(w, "index.html", like)

	w.Write([]byte("Привет из Snippetbox"))
}

func Liker(w http.ResponseWriter,r *http.Request){
	if r.Method != http.MethodPost{
		return
	} else {
		like++
	}
	
}

func showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)

		return
	}

	w.Write([]byte("Создание новой заметки..."))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/like",Liker)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
