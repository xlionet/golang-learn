package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
		{"Name": "Ed", "Text": "Knock knock."}
		{"Name": "Sam", "Text": "Who's there?"}
		{"Name": "Ed", "Text": "Go fmt."}
		{"Name": "Sam", "Text": "Go fmt who?"}
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
	//while the array contains values
	for dec.More() {
		var m Message
		//decode an array value
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}
	//read closing bracket
	t, err = dec.Tocken()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v", t, t)
}

