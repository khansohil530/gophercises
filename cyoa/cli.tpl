{{.Title}}
{{range .Description}}
{{.}}
{{end}}

{{range $i, $v := .Options}}
{{$i}}: {{$v.Text}}
{{end}}