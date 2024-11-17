package tab

const (
	TabName = "tab"
)

type Tab struct{}

func NewTab() *Tab {
	return &Tab{}
}

func (s *Tab) GetCommands() map[string]string {
	return map[string]string{
		":tabnew (or :tabnew {page.words.file})": "open a file in a new tab",
		"Ctrl+wT":                                "move the current split window into its own tab",
		"gt or :tabn[ext]":                       "move to the next tab",
		"gT or :tabp[revious]":                   "move to the previous tab",
		"#gt":                                    "move to the tab number #",
		":tabm[ove] #":                           "move the current tab to the #th position",
		":tabc[lose]":                            "close the current tab and all its windows",
		":tabo[nly]":                             "close all tabs except for the current one",
		":tabdo":                                 "run the command on all tabs (eg: `:tabdo q - close all opened tabs`)",
	}
}
