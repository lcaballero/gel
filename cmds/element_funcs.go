package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var normalTags = []string{
	"A",
	"Abbr",
	"Address",
	"Article",
	"Aside",
	"Audio",
	"B",
	"Bdi",
	"Bdo",
	"Blockquote",
	"Body",
	"Button",
	"Canvas",
	"Caption",
	"Cite",
	"Code",
	"Colgroup",
	"Data",
	"Datalist",
	"Dd",
	"Del",
	"Dfn",
	"Div",
	"Dl",
	"Dt",
	"Em",
	"Fieldset",
	"Figcaption",
	"Figure",
	"Footer",
	"Form",
	"H1",
	"H2",
	"H3",
	"H4",
	"H5",
	"H6",
	"Head",
	"Header",
	"Html",
	"I",
	"Iframe",
	"Ins",
	"Kbd",
	"Label",
	"Legend",
	"Li",
	"Main",
	"Map",
	"Mark",
	"Meter",
	"Nav",
	"Noscript",
	"Object",
	"Ol",
	"Optgroup",
	"Option",
	"Output",
	"P",
	"Pre",
	"Progress",
	"Q",
	"Rb",
	"Rp",
	"Rt",
	"Rtc",
	"Ruby",
	"S",
	"Samp",
	"Script",
	"Section",
	"Select",
	"Small",
	"Span",
	"Strong",
	"Style",
	"Sub",
	"Sup",
	"Table",
	"Tbody",
	"Td",
	"Template",
	"Textarea",
	"Tfoot",
	"Th",
	"Thead",
	"Time",
	"Title",
	"Tr",
	"U",
	"Ul",
	"Var",
	"Video",
}

var selfClosingTags = []string{
	"Area",
	"Base",
	"Br",
	"Col",
	"Embed",
	"Hr",
	"Img",
	"Input",
	"Keygen",
	"Link",
	"Meta",
	"Param",
	"Source",
	"Track",
	"Wbr",
}

func main() {
	buf := bytes.NewBufferString("")
	buf.WriteString(`package gel

var (
// Normal tags requiring closing tag.
`)
	for _, s := range normalTags {
		buf.WriteString(
			fmt.Sprintf(`%s Tag = El("%s", false)
`, s, strings.ToLower(s)),
		)
	}

	buf.WriteString(`

// Void elements must be self closing
`)

	for _, s := range selfClosingTags {
		buf.WriteString(
			fmt.Sprintf(`%s Tag = El("%s", true)
`, s, strings.ToLower(s)),
		)
	}

	buf.WriteString(`
)
`)
	if len(os.Args) > 1 {
		name := fmt.Sprintf("%s.go", os.Args[1])
		err := ioutil.WriteFile(name, buf.Bytes(), 0644)
		if err != nil {
			panic(err)
		}
		return
	}
	fmt.Println(buf.String())
}
