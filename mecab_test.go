package mecab

import (
	"testing"
)

func TestParse(t *testing.T) {
	model := NewModel("")
	if model == nil {
		t.Errorf("model is not created.")
	}
	tagger := model.NewTagger()
	if tagger == nil {
		t.Errorf("tagger is not created.")
	}
	sentence := "すもももももももものうち"
	res := tagger.Parse(sentence)
	if res == "" {
		t.Errorf("no parse result")
	}
}
