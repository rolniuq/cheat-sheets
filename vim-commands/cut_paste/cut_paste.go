package cutpaste

const (
	CutPasteName = "cut_paste"
)

type CutPaste struct{}

func NewCutPaste() *CutPaste {
	return &CutPaste{}
}

func (s *CutPaste) GetCommands() map[string]string {
	return map[string]string{
		"yy":              "yank copy a line",
		"2yy":             "yank copy 2 lines",
		"yw":              "yank copy the characters of the word from the cursor position to the start of the next word",
		"yiw":             "yank copy word under the cursor",
		"yaw":             "yank copy word under the cursor and the space after or before it",
		"y$ or Y":         "yank copy to the end of line",
		"p":               "paste the clipboard after cursor",
		"P":               "paste the clipboard before cursor",
		"gp":              "paste the clipboard after cursor and leave cursor after the new text",
		"gP":              "paste the clipboard before cursor and leave cursor after the new text",
		"dd":              "cut a line",
		"2dd":             "cut 2 lines",
		"dw":              "cut the characters of the word from the cursor position to the start of the next word",
		"diw":             "cut word under the cursor",
		"daw":             "cut word under the cursor and the space after or before it",
		":3,5d":           "delete line from 3 to 5",
		":.,$d":           "delete from the current line to the end on the file",
		":.,1d":           "delete from the current line to beginning of the file",
		":10,1d":          "delete from the 10th line to beginning of the file",
		":g/{pattern}/d":  "delete all lines containing pattern",
		":g!/{pattern}/d": "delete all lines not containing pattern",
		"d$ or D":         "cut to the end of the line",
		"x":               "cut the character",
	}
}
