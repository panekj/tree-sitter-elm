package tree_sitter_elm_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-elm"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_elm.Language())
	if language == nil {
		t.Errorf("Error loading Elm grammar")
	}
}
