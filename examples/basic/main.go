package main

import (
	"fmt"

	. "github.com/lcaballero/gel"
)

func main() {
	DemoDivContainer()
	fmt.Printf("\n\n")
	DemoDoc()
	fmt.Printf("\n\n")
	DemoFunctionalComponent()
}

func DemoDivContainer() {
	Div.Class("container").Text("Hello").ToNode().Println()
}

func DemoDoc() {
	Html5(
		Head(
			Title.Text("Demo HTML Dco"),
		),
		Body(
			H1.Text("Hello, World!"),
		),
	).ToNode().Println()
}

func BreadCrumbComponent(items ...string) View {
	f := Frag()
	for _, s := range items {
		f = f.Add(
			A.Atts("href", "/"+s).Text(s),
			Text(" , "),
		)
	}
	return f
}

func DemoFunctionalComponent() {
	Html5(
		Head(
			Title.Text("Bread Crumb Demo"),
		),
		Body(
			H1.Text("Bread Crumb Demo"),
			BreadCrumbComponent("profile", "settings", "avatar"),
		),
	).ToNode().Println()
}
