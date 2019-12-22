package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trie(t *testing.T) {
	assert := assert.New(t)

	trie := NewTrie()

	trie.Insert("apple")
	assert.True(trie.Search("apple"), "Search apple")
	assert.False(trie.Search("app"), "Search app")
	assert.True(trie.StartsWith("app"), "StartsWith app")

	trie.Insert("app")
	assert.True(trie.Search("app"), "Search app")
}
