package searchreplace

const (
	SearchReplaceName = "search_replace"
)

type SearchReplace struct{}

func NewSearchReplace() *SearchReplace {
	return &SearchReplace{}
}

func (s *SearchReplace) GetCommands() map[string]string {
	return map[string]string{
		"/pattern":       "search for pattern",
		"?pattern":       "search backward for pattern",
		"\vpattern":      "very magic pattern - none alphanumeric characters are interpreted as special regex symbol",
		"n":              "repeat search in the same dir",
		"N":              "repeat search in opposite dir",
		":%s/old/new/g":  "replace all old with new throughout file",
		":%s/old/new/gc": "replace all old with new throughout file with comfirmations",
		":noh[lsearch]":  "remove highlighting of search matches",
	}
}
