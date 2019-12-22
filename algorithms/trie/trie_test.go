package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trie(t *testing.T) {
	assert := assert.New(t)

	trie := NewTrie()
	trie.Insert("abc")

	// Search
	assert.False(trie.Search("ab"), "Search ab in [abc]")
	assert.False(trie.Search("abcd"), "Search abcd in [abc]")
	assert.True(trie.Search("abc"), "Search abc in [abc]")
	// StartsWith
	assert.True(trie.StartsWith("abc"), "Search startWith abc in [abc]")
	assert.True(trie.StartsWith("a"), "Search startWith a in [abc]")
	assert.False(trie.StartsWith("xyz"), "Search startWith xyz in [abc]")

	// ["xyz"]
	trie.Insert("xyz")

	assert.True(trie.Search("xyz"), "Search xyz")
	assert.False(trie.Search("abcxyz"), "Search abcxyz")
	assert.True(trie.StartsWith("xy"), "StartsWith xy")
	assert.False(trie.StartsWith("abcxyz"), "StartsWith abcxyz")
}
