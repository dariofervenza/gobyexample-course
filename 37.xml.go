package main

import (
	"encoding/xml"
	"fmt"
)

// go supports xml and xml like formats using "encoding/xml"

// we cna define a map from each field to an xml tag
type Plant struct {
	XMLName xml.Name `xml:"plant"`   // XMLName element defines the name of the XML element
	Id      int      `xml:"id,attr"` // this means that the field id is an XML attribute rathen than a nested element
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

// we implement the stringify interface
func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	// define an xml representation
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// use MarshalIndent to produce a human readable output
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println("go to xml encode:\n", string(out))
	// got:
	//   <plant id="27">
	//    <name>Coffee</name>
	//      <origin>Ethiopia</origin>
	//      <origin>Brazil</origin>
	//   </plant>

	// add a generic xml header
	fmt.Println(xml.Header + string(out)) // <?xml version="1.0" encoding="UTF-8"?>

	// ------ DECODE XML TO A GO STRUCT ------
	// define a variable to write in
	var p Plant
	// if an error is produced, (cannot decode) the err will have a descriptive msg
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	// create a new example, a different plant
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		// nest all plants under <parent><child>
		Plants []*Plant `xml:"parent>child>plant"`
	}

	// use our 2 exmaples to put under parent child schema
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	// encode nested parent child
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println("\nParent -> child:\n", string(out))
	// <nesting>
	// <parent>
	// 	<child>
	// 	<plant id="27">
	// 		<name>Coffee</name>
	// 		<origin>Ethiopia</origin>
	// 		<origin>Brazil</origin>
	// 	</plant>
	// 	<plant id="81">
	// 		<name>Tomato</name>
	// 		<origin>Mexico</origin>
	// 		<origin>California</origin>
	// 	</plant>
	// 	</child>
	// </parent>
	// </nesting>
}
