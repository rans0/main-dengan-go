package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
)

const form = `
<html>
<body>
	<form action="#" method="post" name="bar">
		<input type="text" name="in"/>
		<input type="submit" value="Submit"/>
	</form>
</html>
</body>
`
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprint(w, "Hello" + req.URL.Path[1:])
}

// handle a simple get request
func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>Hello, world</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	switch request.Method {
	case "GET":
		// display from to the user
		io.WriteString(w, form);
	case "POST":
		/* handle the form data, note that ParseForm must
		   be called before we can extract form data with Form */
		io.WriteString(w, request.FormValue("in"))
	}
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/simple", SimpleServer)
	http.HandleFunc("/form", FormServer)
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
