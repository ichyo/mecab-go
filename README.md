# mecab-go
go wrapper for mecab

# Requirements
MeCab (>= 0.99)

# Example
```go
package main

import (
	"../."
	"fmt"
)

func main() {
	arg := "-O wakati"
	sentence := "太郎はこの本を二郎を見た女性に渡した。"

	model := mecab.NewModel(arg)
	tagger := model.NewTagger()

	result := tagger.Parse(sentence)

	fmt.Println(result)
}
```
