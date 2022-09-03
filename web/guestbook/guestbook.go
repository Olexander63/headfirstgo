package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Guestbook - structure used when displaying view.html.
type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

// check causes log.Fatal for any non-nil errors.
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// viewHandler reads guestbook entries and displays them
// along with the record count.
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}

// newHandler displays a form for entering a record.
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

// createHandler receives a POST request with an entry to add
// and appends it to signatures.
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

// getStrings returns the string segment read from fileName,
// one line for each line of the file.
func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

//func executeTemplate(text string, data interface{}) {
//	tmpl, err := template.New("test").Parse(text)
//	check(err)
//	err = tmpl.Execute(os.Stdout, data)
//	check(err)
//}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)

	//text := "Here's my template!\n"
	//tmpl, err := template.New("test").Parse(text)
	//check(err)
	//err = tmpl.Execute(os.Stdout, nil)
	//check(err)

	//templateText := "Template start\nAction: {{.}}\nTemplate end\n"
	//tmpl, err := template.New("test").Parse(templateText)
	//check(err)
	//err = tmpl.Execute(os.Stdout, "ABC")
	//check(err)
	//err = tmpl.Execute(os.Stdout, 42)
	//check(err)
	//err = tmpl.Execute(os.Stdout, true)
	//check(err)

	//executeTemplate("Dot is : {{.}}!\n", "ABC")
	//executeTemplate("Dot is: {{.}}!\n", 123.5)

	//executeTemplate("start {{if.}}Dot ist true!{{end}} finish\n", true)
	//executeTemplate("start {{if.}}Dot ist true!{{end}} finish\n", false)

	//templateText := "Before loop:{{.}}\n{{range .}}In loop {{.}}\n {{end}}After loop: {{.}}=\n"
	//executeTemplate(templateText, []string{"do", "re", "mi"})
	//
	//templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	//executeTemplate(templateText, []float64{1.25, 0, 99, 27})
	//
	//templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	//executeTemplate(templateText, []float64{})
	//executeTemplate(templateText, nil)

	//templateText := "Name:{{.Name}}\nCount:{{.Count}}\n"
	//executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
	//executeTemplate(templateText, Part{Name: "Cables", Count: 2})

	//templateText := "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	//subscriber := Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	//executeTemplate(templateText, subscriber)
	//subscriber = Subscriber{Name: "Joy Carr", Rate: 5.99, Active: false}
	//executeTemplate(templateText, subscriber)
}
