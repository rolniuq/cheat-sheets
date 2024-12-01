package help

const (
	HelpName = "help"
)

type Help struct{}

func NewHelp() *Help {
	return &Help{}
}

func (h *Help) GetCommands() map[string]string {
	return map[string]string{
		"cursor":         "show cursor commands",
		"cut_paste":      "show cut and paste commands",
		"edit":           "show edit commands",
		"exiting":        "show exit commands",
		"global":         "show global commands",
		"indent_text":    "show indent text commands",
		"insert":         "show insert mode commands",
		"macro":          "show macro commands",
		"make_text":      "show make text commands",
		"mark_positon":   "show mark position commands",
		"register":       "show register commands",
		"search_mul":     "show search multiple commands",
		"search_replace": "show search and replace commands",
		"tab":            "show tab commands",
		"visual_cmd":     "show visual commands",
	}
}
