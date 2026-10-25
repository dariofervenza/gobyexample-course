package main

import (
	"bytes"
	"encoding/json/v2"
	"fmt"
	"strings"
)

// go includes JSON support, from a data type to json or from json to a data type

type response1 struct {
	Page   int
	Fruits []string
}

// DISCLAIMER:
// only exported fields will be visible to json lib (capitalized)
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// ENCODE BASIC TYPES TO JSON
	bolB, _ := json.Marshal(true)
	fmt.Println("Encoded json bool: ", string(bolB))

	// ENCODE A INTEGER
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	// ENCODE A FLOAT
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	// ENCODE A STRING
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// ENCODE A SLICE
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("slice: ", string(slcB))

	// ENCODE A MAP
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println("map: ", string(mapB))

	// ------------- custom dtypes --------------------
	// IF NO CUSTOM TAGS --> JSON FIELDS WILL BE EXACTLY LIKE FIELDNAMES, only exported fields included
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("no tag struct: ", string(res1B)) // {"Page":1,"Fruits":["apple","peach","pear"]} (upercase keys)

	// if struct has tags, json will use tag names instead of fieldNames
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("tagged struct: ", string(res2B))

	// --------------- decode JSON to go objects -----------------------
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// json decode works inplace, hence it needs a variable to write in
	var dat map[string]any
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println("decoded map: ", dat)

	// convert num to float64
	num := dat["num"].(float64)
	fmt.Println("num in float64: ", num)

	// to access a series of nested data we need some convertions
	strs := dat["strs"].([]any)
	str1 := strs[0].(string)
	fmt.Println("elemet 0 in nested str", str1)

	// ------ DECODE JSON INTO A STRUCT  -----
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	_ = json.Unmarshal([]byte(str), &res)
	fmt.Println("STRUCT DECODED: ", res)
	fmt.Println("decode to go object: ", res.Fruits[0])

	// ----- STREAM JSON ENCODING INTO A BUFFER --> io.Writer (or http response bodies, or os.Stdout)
	// MAP -> IO BUFFER
	d := map[string]int{"apple": 5, "lettuce": 7}
	var buf bytes.Buffer
	_ = json.MarshalWrite(&buf, d)
	fmt.Println("to io buffer: ", buf.String())

	// read from streaming io like os.Stdin, io.Readers or HTTP request bodies
	res1 := response2{}
	_ = json.UnmarshalRead(strings.NewReader(str), &res1) // fake a reader
	fmt.Println("from io: ", res1)

	// check json and go blog: https://go.dev/blog/json

	// check json lib docs: https://pkg.go.dev/encoding/json/v2
}
