package edit

const (
	EditName = "edit"
)

type Edit struct{}

func NewEdit() *Edit {
	return &Edit{}
}

func (s *Edit) GetCommands() map[string]string {
	return map[string]string{
		"r":        "replace a single character",
		"R":        "replace more than one character, util ESC pressed",
		"J":        "join line below to current one with one space in between",
		"gJ":       "join line below to current one without space in between",
		"gwip":     "reflow paragraph",
		"g~":       "switch case up to motion",
		"gu":       "change to lowercase up to motion",
		"gU":       "change to uppercase up to motion",
		"cc":       "change entire line",
		"c$ or C":  "replace to the end of the line",
		"ciw":      "change entire word",
		"cw or ce": "change to the end of word",
		"s":        "delete character and substitute text (same as cl)",
		"S":        "delete line and substitute text (same as cc)",
		"xp":       "transpose two letters (delete and paste)",
		"u":        "undo",
		"U":        "restore last changed line",
		"Ctrl+u":   "redo",
		".":        "repeat last command",
		"~":        "Changes the case of current character",
		"guu":      "Change current line from upper to lower",
		"gUU":      "Change current LINE from lower to upper",
		"guw":      "Change to end of current WORD from upper to lower",
		"guaw":     "Change all of current WORD to lower",
		"gUw":      "Change to end of current WORD from lower to upper",
		"gUaw":     "Change all of current WORD to upper",
		"g~~":      "Invert case to entire line",
		"g~w":      "Invert case to current WORD",
		"guG":      "Change to lowercase until the end of document",
		"gU)":      "Change until end of sentence to upper case",
		"gu}":      "Change to end of paragraph to lower case",
		"gU5j":     "Change 5 lines below to upper case",
		"gu3k":     "Change 3 lines above to lower case",
	}
}
