package main

import "fmt"

// go supports embedding of structs and interfaces

// DO NOT CONFUSE them with go:embed (https://gobyexample.com/embed-directive)
// the later is a directive introduced in go 1.16
// and its used for embed files and folders into the application binary

// a container embeds a base. an embedding looks like a field without a name

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	// when creating structs with literals
	// we need to initialize the embedding explicitly
	// here the embedded type serves as the field name
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// now we can access directly the base attr without going through .base
	// eg. co.num instead of c.base.num
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// absolute path can also be used
	fmt.Println("Directly: ", co.base.num)

	// methods of base also are added to container
	// but its not inheritance, you cannot overwirte it with func (c container) describe() string {} --> even adding this would call base describe()
	fmt.Println("describe: ", co.describe())

	// add embeddings into an interface:

	type describer interface {
		describe() string
	}

	// we see that a container implements describer by assigning its type
	var d describer = co
	fmt.Println("With the interface: ", d.describe())

}
