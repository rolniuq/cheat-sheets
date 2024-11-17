package maketext

const (
	MakeTextName = "make_text"
)

type MakeText struct{}

func NewMakeText() *MakeText {
	return &MakeText{}
}

func (s *MakeText) GetCommands() map[string]string {
	return map[string]string{
		"v":             "start visual mode, mark lines, then do a command (ex: yank)",
		"V":             "start linewise visual mode",
		"o":             "move to the end of marked area",
		"Ctrl+v":        "start visual block mode",
		"0":             "move to other corner of block",
		"aw":            "mark a word",
		"ab":            "a block with ()",
		"at":            "a block with <> tags",
		"ib":            "inner block with ()",
		"iB":            "inner block with {}",
		"it":            "inner block with <> tags",
		"ESC or Ctrl+c": "exit visual mode",
		"Tips":          "b or B can also use ( or { respectively",
	}
}
