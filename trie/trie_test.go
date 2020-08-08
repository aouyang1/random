package main

import "testing"

func TestInsert(t *testing.T) {
	tr := NewTrie()

	// try to add an empty word to a new trie
	tr.Insert("")
	if tr.NumWords() != 0 {
		t.Errorf("Expected no words in the trie while adding an empty string, but got %d\n", tr.NumWords())
		return
	}

	// add a word to new trie
	tr.Insert("asdf")
	if tr.NumWords() != 1 {
		t.Errorf("Expected 1 word on new word entry but got %d\n", tr.NumWords())
		return
	}

	// add a duplicate
	tr.Insert("asdf")
	if tr.NumWords() != 1 {
		t.Errorf("Expected 1 word after duplicate entry, but got %d\n", tr.NumWords())
		return
	}

	// add a new unique word
	tr.Insert("as")
	if tr.NumWords() != 2 {
		t.Errorf("Expected 2 words after an additional unique word but got %d\n", tr.NumWords())
		return
	}
}

func TestSearch(t *testing.T) {
	tr := NewTrie()
	tr.Insert("")
	tr.Insert("asdf")
	tr.Insert("as")
	tr.Insert("qu")
	tr.Insert("query")
	tr.Insert("queries")

	testData := []struct {
		query    string
		expected []string
	}{
		{"as", []string{"as", "asdf"}},
		{"a", []string{"as", "asdf"}},
		{"asd", []string{"asdf"}},
		{"asdfg", []string{}},
		{"quer", []string{"query", "queries"}},
	}
	for _, td := range testData {
		res := tr.Search(td.query)
		if len(res) != len(td.expected) {
			t.Errorf("Expected %d results, but got %d, %v\n", len(td.expected), len(res), res)
			return
		}

		var found bool
		for _, r := range res {
			found = false
			for _, er := range td.expected {
				if er == r {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Did not find %s in expected\n", r)
				return
			}
		}
		for _, er := range td.expected {
			found = false
			for _, r := range res {
				if er == r {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Did not find %s in result\n", er)
				return
			}
		}
	}
}

func TestRemove(t *testing.T) {
	tr := NewTrie()
	tr.Insert("")
	tr.Insert("asdf")
	tr.Insert("as")
	tr.Insert("qu")
	tr.Insert("query")
	tr.Insert("queries")

	testData := []struct {
		word     string
		expected bool
		numWords int
	}{
		{"query", true, 4},
		{"que", false, 4},
		{"asdf", true, 3},
	}
	for _, td := range testData {
		removed := tr.Remove(td.word)
		if removed != td.expected {
			t.Errorf("Expected removed to be %t for word, %s\n", td.expected, td.word)
			return
		}
		if tr.NumWords() != td.numWords {
			t.Errorf("Expected %d remaining words, but got %d\n", td.numWords, tr.NumWords())
			return
		}
	}
}
