package main

import (
	"fmt"
	"text/template"
	"os"
)

var allTemplate = `
{{ define "a" }} Template A {{ end }}
{{ define "b" }} Template B {{ end }}
{{ define "all" }} {{template "a"}} {{template "b"}} {{ end }}
`

func main() {
	t := template.New("allTemplate")
	t, err := t.Parse(allTemplate)
	if err != nil {
		fmt.Println("parsing failed: %s", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "a", nil)
	if err != nil {
		fmt.Println("execution failed: %s", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "b", nil)
	if err != nil {
		fmt.Println("execution failed: %s", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "all", nil)
	if err != nil {
		fmt.Println("execution failed: %s", nil)
	}
}
