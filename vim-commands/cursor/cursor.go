package cursor

const (
	CursorName = "cursor"
)

type Cursor struct{}

func NewCursor() *Cursor {
	return &Cursor{}
}

func (c *Cursor) GetCommands() map[string]string {
	return map[string]string{
		"j":         "move cursor down",
		"k":         "move cursor up",
		"h":         "move cursor left",
		"l":         "move cursor right",
		"gj":        "move cursor down by line",
		"gk":        "move cursor up by line",
		"H":         "move cursor to top of screen",
		"M":         "move cursor to middle of screen",
		"L":         "move cursor to bottom of screen",
		"w":         "jump cursor to the start of a word",
		"W":         "jump cursor to the start of a word (word can contain punctuation)",
		"e":         "jump cursor to the end of a word",
		"E":         "jump cursor to the end of a word (word can contain punctuation)",
		"b":         "jump cursor back to the start of a word",
		"B":         "jump cursor back to the start of a word (word can contain punctuation)",
		"ge":        "jump cursor back to the end of a word",
		"gE":        "jump cursor back to the end of a word (word can contain punctuation)",
		"%":         "jump cursor to matching character ('{}', '[]', '()')",
		"0":         "jump cursor to the start of the line",
		"^":         "jump cursor to the first non-blank character of the line",
		"$":         "jump cursor to the end of the line",
		"g_":        "jump cursor to the last non-blank character of the line",
		"gg":        "jump cursor to the first line of document",
		"G":         "jump cursor to the last line of document",
		"5gg or 5G": "jump cursor to the 5's line of document",
		"gd":        "jump to local definition",
		"gD":        "jump to global definition",
		"fx":        "jump to the next occurrence of x",
		"tx":        "jump to before next occurrence of x",
		"Fx":        "jump to the previous occurrence of x",
		"Tx":        "jump to after previous occurrence of x",
		";":         "repeat previous f, t, F or T movement",
		",":         "repeat previous f, t, F or T movement backward",
		"}":         "jump to next paragraph",
		"{":         "jump to previous paragraph",
		"zz":        "center cursor on screen",
		"zt":        "position cursor at top of screen",
		"zb":        "position cursor at bottom of screen",
		"Ctrl+e":    "move screen down 1 line",
		"Ctrl+y":    "move screen up 1 line",
		"Ctrl+b":    "move screen up one page",
		"Ctrl+f":    "move screen down one page",
		"Ctrl+d":    "move cursor/screen down half a page",
		"Ctrl+u":    "move cursor/screen up half a page",
	}
}
