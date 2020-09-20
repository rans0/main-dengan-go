package main

import (
	"fmt"
	"net/http"
	"html/template"
)

const port = 3000
var guestList []string

var indexHTML = `
<!DOCTYPE html>
<html>
    <head>
		<title>Guest Book ::Web GUI</title>
    </head>
    <body>
		<h1>Guest Book :: Web GUI</h1>
		<form action="/add" method="post">
		Name: <input name="name" /><submit value="Sign Guest Book">
		</form>
		<hr />
		<h4>Previous Guests</h4>
		<ul>
			{{range .}}
			<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
</html>
`

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	t, err := t.Parse(indexHTML)
	if err != nil {
		message := fmt.Sprintf("Bad template : %s", err)
		http.Error(w, message, http.StatusInternalServerError)
	}
	t.Execute(w, guestList)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	guest := r.FormValue("name")
	if len(guest) > 0 {
		guestList = append(guestList, guest)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}