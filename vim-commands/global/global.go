package global

const (
	GlobalName = "global"
)

type Global struct{}

func NewGlobal() *Global {
	return &Global{}
}

func (g *Global) GetCommands() map[string]string {
	return map[string]string{
		":h[elp]":     "open for help keyword",
		":sav[eas]":   "save file as",
		":clo[se]":    "close current pane",
		":ter[minal]": "open a terminal window",
		"K":           "open man page for word under cursor",
	}
}
