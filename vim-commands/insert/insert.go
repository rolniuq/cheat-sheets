package insert

const (
	InsertName = "insert"
)

type Insert struct{}

func NewInsert() *Insert {
	return &Insert{}
}

func (i *Insert) GetCommands() map[string]string {
	return map[string]string{
		"i":       "insert before the cursor",
		"I":       "insert at the start of the line",
		"a":       "append at after cursor",
		"A":       "append at the end of the line",
		"o":       "append a new line below the current line",
		"O":       "append a new line above the current line",
		"ea":      "append at the end of word",
		"Ctrl+h":  "delete a character before the cursor during insert mode",
		"Ctrl+w":  "delete a word before the cursor during insert mode",
		"Ctrl+j":  "add a line break at the cursor position during insert mode",
		"Ctrl+t":  "move right line one shift-width during insert mode",
		"Ctrl+d":  "move left line one shift-width during insert mode",
		"Ctrl+n":  "auto complete next match before the cursor during insert mode",
		"Ctrl+p":  "auto complete previous match before the cursor during insert mode",
		"Ctrl+rx": "insert the contents of register x",
		"Ctrl+ox": "temporarily enter normal mode to issue one normal mode command x",
	}
}
