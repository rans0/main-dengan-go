package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"sort"
	"strconv"
	"strings"
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

const formCalculate = `
<html>
<body>
<form action="/calculate" method="POST">
	<h1>Statistics</h1>
	<h5>Compute base statistics for a given list of numbers</h5>
	<label for="numbers">Numbers (comma or space-separated):</label><br>
	<input type="text" name="numbers" size="30"><br />
	<input type="submit" value="Calculate">
</form>
</html>
</body>
`

const error = `<p class="error">%s</p>`

var pageTop, pageBottom = "", ""

type statistics struct {
	numbers []float64
	mean float64
	median float64
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprint(w, "Hello" + req.URL.Path[1:])
}

// handle a simple get request
func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>Hello, world</h1>")
}

// Write an HTML header, parse the form, write form to writer and make request for numbers
func CalculateServer(w http.ResponseWriter, request *http.Request) { // Capture the numbers from the request, and format the data and check for error
	w.Header().Set("Content-Type", "text/html")

	err := request.ParseForm() // Must be called before writing response
	fmt.Fprint(w, pageTop, formCalculate)
	if err != nil {
		fmt.Fprint(w, error, err)
	} else {
		if numbers, message, ok := processRequest(request); ok {
			stats := getStats(numbers)
			fmt.Fprint(w, formatStats(stats))
		} else if message != "" {
			fmt.Fprint(w, error, err)
		}
	}
	fmt.Println("calculate server running")
	fmt.Fprint(w, pageBottom)
}

// Capture the numbers from the request, and format the data and check for error
func processRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64

	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers,"'" + field + "' is invalid", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false // no data first time form is shown
	}
	return numbers, "", true
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return
}

func median(numbers []float64) (median float64) {
	mid := len(numbers) / 2
	result := numbers[mid]
	if len(numbers) % 2 == 0 {
		result = (result + numbers[mid-1]) / 2
	}
	return result
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	switch request.Method {
	case "GET":
		// display form to the user
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
	http.HandleFunc("/calculate", CalculateServer)
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
