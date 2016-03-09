# Introduction

`gel` is lib for programmatically producing Html.

## Usage

```go

package main
import (
  . "github.com/lcaballero/gel"
  "fmt"
)  

func main() {
  el := Div.Add(
    Att("class", "container"),
    Att("id", "id-1"),
    Text("text"),
  )
  html := el.String() // <div class="container", id="id-1">text</div>
  fmt.Println(html)  
}

```


## License

See license file.

The use and distribution terms for this software are covered by the
[Eclipse Public License 1.0][EPL-1], which can be found in the file 'license' at the
root of this distribution. By using this software in any fashion, you are
agreeing to be bound by the terms of this license. You must not remove this
notice, or any other, from this software.


[EPL-1]: http://opensource.org/licenses/eclipse-1.0.txt
