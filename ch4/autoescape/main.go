// Autoescape demonstrates automatic HTML escaping in html/template.
// Autoescape 演示了在 html/template 中的自动 HTML 转义。
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string        // untrusted plain text // 不受信任的纯文本
		B template.HTML // trusted HTML 	//受信任的html
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
