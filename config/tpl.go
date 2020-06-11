package config

import "html/template"

//Tpl to hold all templates
var Tpl *template.Template

func init() {
	//template
	Tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}
