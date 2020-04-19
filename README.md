# 🍌split [![](https://godoc.org/github.com/tachiniererin/bananasplit?status.svg)](https://godoc.org/github.com/tachiniererin/bananasplit)

Bananasplit is a simple library to split strings by Unicode code-point ranges.

## Example
```go
package main

import (
	"fmt"

	"github.com/tachiniererin/bananasplit"
)

func main() {
	// define only the emoji range, to split by emoji sequences
	// and everything else
	r := map[string][]bananasplit.RuneRange{
		"emoji": bananasplit.EmojiRange,
	}
	s := "tachiniererin🏳️‍🌈🏳️‍⚧️ emoji splitting"
	fmt.Printf("%+v", bananasplit.SplitByRanges(s, r))
}
```
