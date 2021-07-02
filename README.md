# Library to resolve the picking problem

## Usage 
```
go get github.com/DarioRoman01/picking_problem
```

```go
package main

import (
    "fmt"
    p "github.com/DarioRoman01/picking_problem"
)

func main() {
    recommendations := p.ParseFile("./path/to/your/file.json")
    tokens := FindRecomendations(recommendations)
    fmt.Println(tokens)
}
```