package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type EmailData struct {
	RecieptName string
	SenderName  string
	Subject     string
	Body        string
	Items       []string // demo a loop
	UnreadCount int
}

func main() {
	fmt.Println("------ Text Template Example -----")

	emailTemplate := `
Subject : {{ .Subject }}
{{ .Body }}
{{ if .Items }}
	Related Items: 
{{ range .Items }}
	- {{.}}
{{end}}
{{end}}

{{ if gt .UnreadCount 0 }}
You have {{.UnreadCount}} unreads.
{{ else }}
You have no messages
{{ end }}

- Thanks
{{.SenderName}}
`

	tmpl, err := template.New("email-message").Parse(emailTemplate)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data := EmailData{
		RecieptName: "Alice",
		SenderName:  "Bob's Auto responder",
		Subject:     "Your weekly update",
		Body:        "Here is the update you requested. hope its useful",
		Items:       []string{"Report A", "Document B"},
		UnreadCount: 0,
	}

	var output strings.Builder

	err = tmpl.Execute(&output, data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(strings.ToUpper(output.String()))
}
