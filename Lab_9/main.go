package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"html/template"
	"net/http"
)

func main() {
	myHandler := http.HandlerFunc(myFunc)
	fmt.Println("Listening on http://127.0.0.1:8090/")
	http.ListenAndServe(":8090", nosurf.New(myHandler))
}

var templ = template.Must(template.New("t1").Parse(templateString))

func myFunc(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]string)
	context["token"] = nosurf.Token(r)
	if r.Method == "POST" {
		context["name"] = r.FormValue("name")
	}

	templ.Execute(w, context)
}



var templateString = `
<!doctype html>
<html>
<body style="background-color:blue">
{{ if .name }}
<p style="background-color:yellow;  font-family: courier;" >Your name: {{ .name }}</p>
{{ end }}
<form action="/" method="POST">
<input type="text" name="name">

<!-- Try removing this or changing its value
     and see what happens -->
<input type="hidden" name="csrf_token" value="{{ .token }}">
<input type="submit" value="Send">
</form>
</body>
</html>
`