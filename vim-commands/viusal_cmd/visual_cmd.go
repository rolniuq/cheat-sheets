package viusalcmd

const (
	VisualCmdName = "visual_cmd"
)

type VisualCmd struct{}

func NewVisualCmd() *VisualCmd {
	return &VisualCmd{}
}

func (s *VisualCmd) GetCommands() map[string]string {
	return map[string]string{
		">": "shift text right",
		"<": "shift text left",
		"y": "copy marked text",
		"d": "delete marked text",
		"~": "switch case",
		"u": "change marked text to lowercase",
		"U": "change marked text to uppercase",
	}
}
