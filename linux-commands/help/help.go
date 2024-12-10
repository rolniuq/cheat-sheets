package help

const HelpName = "help"

type Help struct{}

func NewHelp() *Help {
	return &Help{}
}

func (h *Help) GetCommands() map[string]string {
	return map[string]string{
		"basic": "Get basic commands",
	}
}
