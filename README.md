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
	tokens, err := p.FindMostHomogenousRecommendations("./test.json")
	if err != nil {
		fmt.Println(err.Error())
	}
    
	fmt.Println(tokens)
}
```
