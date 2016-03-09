package gel

type Type int

const (
	Textual   Type = 1
	Element   Type = 2
	Attribute Type = 3
)

//go:generate stringer -type=Tag
type Tag int

const (
	A Tag = iota + 1
	Abbr
	Address
	Article
	Aside
	Audio
	B
	Bdi
	Bdo
	Blockquote
	Body
	Button
	Canvas
	Caption
	Cite
	Code
	Colgroup
	Data
	Datalist
	Dd
	Del
	Dfn
	Div
	Dl
	Dt
	Em
	Fieldset
	Figcaption
	Figure
	Footer
	Form
	H1
	H2
	H3
	H4
	H5
	H6
	Head
	Header
	Html
	I
	Iframe
	Ins
	Kbd
	Label
	Legend
	Li
	Main
	Map
	Mark
	Meter
	Nav
	Noscript
	Object
	Ol
	Optgroup
	Option
	Output
	P
	Pre
	Progress
	Q
	Rb
	Rp
	Rt
	Rtc
	Ruby
	S
	Samp
	Script
	Section
	Select
	Small
	Span
	Strong
	Style
	Sub
	Sup
	Table
	Tbody
	Td
	Template
	Textarea
	Tfoot
	Th
	Thead
	Time
	Title
	Tr
	U
	Ul
	Var
	Video

	// Void element (ie self-closing tags).
	Area
	Base
	Br
	Col
	Embed
	Hr
	Img
	Input
	Keygen
	Link
	Meta
	Param
	Source
	Track
	Wbr
)

var AllTags = []Tag{
	A,Abbr,Address,Article,Aside,Audio,B,Bdi,Bdo,Blockquote,Body,Button,Canvas,
	Caption,Cite,Code,Colgroup,Data,Datalist,Dd,Del,Dfn,Div,Dl,Dt,Em,Fieldset,
	Figcaption,Figure,Footer,Form,H1,H2,H3,H4,H5,H6,Head,Header,Html,I,Iframe,
	Ins,Kbd,Label,Legend,Li,Main,Map,Mark,Meter,Nav,Noscript,Object,Ol,
	Optgroup,Option,Output,P,Pre,Progress,Q,Rb,Rp,Rt,Rtc,Ruby,S,Samp,Script,
	Section,Select,Small,Span,Strong,Style,Sub,Sup,Table,Tbody,Td,Template,
	Textarea,Tfoot,Th,Thead,Time,Title,Tr,U,Ul,Var,Video,

	// Void element (ie self-closing tags).
	Area,Base,Br,Col,Embed,Hr,Img,Input,Keygen,Link,Meta,Param,Source,Track,Wbr,
}

var NormalTags = []Tag{
	A,Abbr,Address,Article,Aside,Audio,B,Bdi,Bdo,Blockquote,Body,Button,Canvas,
	Caption,Cite,Code,Colgroup,Data,Datalist,Dd,Del,Dfn,Div,Dl,Dt,Em,Fieldset,
	Figcaption,Figure,Footer,Form, H1, H2, H3, H4, H5, H6,Head,Header,Html,I,
	Iframe,Ins,Kbd,Label,Legend,Li,Main,Map,Mark,Meter,Nav,Noscript,Object,
	Ol,Optgroup,Option,Output,P,Pre,Progress,Q,Rb,Rp,Rt,Rtc,Ruby,S,Samp,
	Script,Section,Select,Small,Span,Strong,Style,Sub,Sup,Table,Tbody,Td,
	Template,Textarea,Tfoot,Th,Thead,Time,Title,Tr,U,Ul,Var,Video,
}

var VoidTags = []Tag{
	Area,Base,Br,Col,Embed,Hr,Img,Input,Keygen,
	Link,Meta,Param,Source,Track,Wbr,
}

var voidSet map[Tag]struct{} = nil

func IsClosing(tag Tag) bool {
	_,ok := voidSet[tag]
	return ok
}

func init() {
	voidSet = make(map[Tag]struct{})
	for _,v := range VoidTags {
		voidSet[v] = struct{}{}
	}
}

