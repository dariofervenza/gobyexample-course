package main

import (
	"os"
	"text/template"
)

// stdlib: text/templates
// use it to generate dynamic content or show customized info to the user

// NOTICE that there is a similar package html/template --> it provides the same API but with additional security features
// use html/template when generating HTML

func main() {
	t1 := template.New("t1")

	// parse its body from a string
	// templates dynamically insert content in their "ACTIONS"
	// actions are enclosed in {{...}}
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// we can use Must to make it panic in case Parse returns an error --> useful for templates initialized in the global scope
	t1 = template.Must(t1.Parse("with must Value: {{.}}\n"))

	// to apply an input text to a tmeplate, we execute it:
	t1.Execute(os.Stdout, "some random string")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"cpp",
	})

	// define a helper

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// access to a struct field it has to be exported (be public by being Capitalized)
	t2 := Create("t2", "CHECK STRUCT / map field (called Name): {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name  string
		other int
	}{"Jane Doe", 5})

	// but it also works with Maps, in this case it does not require to be capitalized
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse", "otherkey": "6",
	})

	// CONDITIONAL EXECUTION OF TEMPLATES
	// add if else, a vlaue is ocnsidered false if its the zero vlaue of that type (nil pointer, 0, "",  ...)
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n") // prints yes or no
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// ------ TEMPLATES WITH ITERATION ---------
	// use Range to iterabte through slices, arrays, maps ir channels
	// inside the range block {{.}} is set to the current item of the iteration
	// it will print (in this example): Range element1 element2 element3
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
	t5 := Create("t5", "Map: {{range $key, $value := .}}{{$key}} is {{$value}}, {{end}}\n")
	t5.Execute(os.Stdout,
		map[string]string{
			"Go":   "Fast",
			"Rust": "Safe",
			"C++":  "Powerful",
			"C#":   "Versatile",
		})
	t6 := Create("t6", "Languages:\n{{range $i, $v := .}}  {{$i}}: {{$v}}\n{{end}}")
	t6.Execute(os.Stdout,
		[]string{
			"Python",
			"JavaScript",
			"Ruby",
			"Swift",
		})
}
