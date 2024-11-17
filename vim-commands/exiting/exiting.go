package exiting

const (
	ExitingName = "exiting"
)

type Exiting struct{}

func NewExiting() *Exiting {
	return &Exiting{}
}

func (s *Exiting) GetCommands() map[string]string {
	return map[string]string{
		":w":              "save the file but not exit",
		":w !sudo tee %":  "write out the current file using sudo",
		":wq or :x or ZZ": "save and quite",
		":q! or ZQ":       "quit and throw away unsaved changes",
		":wqa":            "save and quit on all tabs",
	}
}
