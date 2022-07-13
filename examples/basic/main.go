package main

import (
	"fmt"
	"os"

	. "github.com/lcaballero/gel"
)

func main() {
	DemoDivContainer()
	fmt.Printf("\n\n")
	DemoDoc()
}

func DemoDivContainer() {
	d := Div.Class("container").Text("Hello")
	d.ToNode().WriteWithIndention(NewIndent(), os.Stdout)
}

func DemoDoc() {
	Html5(
		Head(
			Title.Text("Demo HTML Dco"),
		),
		Body(
			H1.Text("Hello, World!"),
		),
	).ToNode().WriteWithIndention(NewIndent(), os.Stdout)
}
