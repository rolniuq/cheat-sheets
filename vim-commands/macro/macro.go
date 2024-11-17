package macro

const (
	MacroName = "macro"
)

type Macro struct{}

func NewMacro() *Macro {
	return &Macro{}
}

func (s *Macro) GetCommands() map[string]string {
	return map[string]string{
		"qa": "record macro a",
		"q":  "stop recording macro",
		"@a": "run macro a",
		"@@": "return last run macro",
	}
}
