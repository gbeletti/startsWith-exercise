package main

import "testing"

func TestStartsWith(t *testing.T) {
	tr := setupTest()
	tcases := []struct {
		Name     string
		Prefix   string
		Expected []string
	}{
		{
			Name:     "01 - words with ho",
			Prefix:   "ho",
			Expected: []string{"home", "homework", "house", "housekeeper"},
		},
		{
			Name:     "02 - words with car",
			Prefix:   "car",
			Expected: []string{"car", "cars", "card"},
		},
		{
			Name:     "03 - words with card",
			Prefix:   "card",
			Expected: []string{"card"},
		},
		{
			Name:     "04 - no result",
			Prefix:   "zzz",
			Expected: []string{},
		},
		{
			Name:     "05 - words with drin",
			Prefix:   "drink",
			Expected: []string{"drink", "drinking"},
		},
	}
	for _, tcase := range tcases {
		t.Run(tcase.Name, func(t *testing.T) {
			testStartsWith(t, tr, tcase.Prefix, tcase.Expected)
		})
	}
}

func testStartsWith(t *testing.T, tr *trie, prefix string, expected []string) {
	got := tr.StartsWith(prefix)
	mapGot := make(map[string]struct{})
	for _, word := range got {
		mapGot[word] = struct{}{}
	}
	for _, ex := range expected {
		if _, ok := mapGot[ex]; !ok {
			t.Errorf("expected [%s] and didnt find", ex)
			continue
		}
		delete(mapGot, ex)
	}
	for word := range mapGot {
		t.Errorf("got [%s] and shouldnt have found it", word)
	}
}

func TestSearch(t *testing.T) {
	tr := setupTest()
	tcases := []struct {
		Name     string
		Search   string
		Expected bool
	}{
		{
			Name:     "01 - search found",
			Search:   "drink",
			Expected: true,
		},
		{
			Name:     "02 - search not found",
			Search:   "drinki",
			Expected: false,
		},
		{
			Name:     "03 - search found",
			Search:   "car",
			Expected: true,
		},
		{
			Name:     "04 - search not found",
			Search:   "appli",
			Expected: false,
		},
		{
			Name:     "05 - search not found",
			Search:   "apple",
			Expected: true,
		},
	}

	for _, tcase := range tcases {
		t.Run(tcase.Name, func(t *testing.T) {
			testSearchWord(t, tr, tcase.Search, tcase.Expected)
		})
	}
}

func testSearchWord(t *testing.T, tr *trie, search string, expected bool) {
	got := tr.SearchWord(search)
	if got != expected {
		t.Errorf("search test failed. expected [%t] got [%t]", expected, got)
	}
}

func setupTest() *trie {
	words := getWords()
	tr := newTrie()
	for _, w := range words {
		tr.Add(w)
	}
	return tr
}

func getWords() (words []string) {
	words = []string{
		"apple",
		"artist",
		"banana",
		"boat",
		"book",
		"car",
		"cars",
		"card",
		"drink",
		"drinking",
		"dumb",
		"home",
		"homework",
		"house",
		"housekeeper",
	}
	return
}
