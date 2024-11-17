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
		"cursor": "show cursor commands",
		"global": "show global commands",
		"insert": "show insert mode commands",
	}
}
