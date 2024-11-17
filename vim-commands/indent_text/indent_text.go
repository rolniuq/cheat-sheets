package indenttext

const (
	IndentTextName = "indent_text"
)

type IndentText struct{}

func NewIndentText() *IndentText {
	return &IndentText{}
}

func (s *IndentText) GetCommands() map[string]string {
	return map[string]string{
		">>":   "move right line one shiftwidth",
		"<<":   "move left line one shiftwidth",
		">%":   "indent a block with () or {} (cursor on brace)",
		"<%":   "de-indent a block with () or {} (cursor on brace)",
		">ib":  "indent inner block with ()",
		">at":  "indent a block with <> tags",
		"3==":  "re-indent 3 lines",
		"=%":   "re-indent a block with () or {} (cursor on brace)",
		"=iB":  "re-indent inner block with {}",
		"gg=G": "re-indent entire buffer",
		"]p":   "paste and adjust indent to current line",
	}
}
