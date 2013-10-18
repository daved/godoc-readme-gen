//Copyright 2013 James Frasche. All rights reserved.
//Use of this code is governed by a BSD-License found in the LICENSE file.

package main

import "text/template"

var tmplraw = `#{{.Name}}
{{.Synopsis}}

Download:
` + "```" + `shell
go get {{.Import}}
` + "```" + `
{{if .Library}}

Full documentation at http://godoc.org/{{.Import}}
{{end}}
* * *
{{.Doc}}
{{if .Bugs}}
#Bugs
{{range .Bugs}}* {{.}}{{end}}
{{end}}
* * *
Automatically generated by [autoreadme](https://github.com/JImmyFrasche/autoreadme) on {{.Today}}
`
var tmpl = template.Must(template.New("").Parse(tmplraw))