package markposition

const (
	MarkPositionName = "mark_position"
)

type MarkPosition struct{}

func NewMarkPosition() *MarkPosition {
	return &MarkPosition{}
}

func (s *MarkPosition) GetCommands() map[string]string {
	return map[string]string{
		":marks":   "list of marks",
		"ma":       "set current position of mark A",
		"`a":       "jump to position of mark A",
		"y`a":      "yank text to position of mark A",
		"`0":       "go to the position where Vim was previously exited",
		"`''":      "go to the position when last editing this file",
		"`.":       "go to the position of the last change in this file",
		"``":       "go to the position before last jump",
		":ju[mps]": "list of jumps",
		"Ctrl+i":   "go to newer position in jump list",
		"Ctrl+o":   "go to older position in jump list",
		":changes": "list of changes",
		"g,":       "go to newer position in change list",
		"g;":       "go to older position in change list",
		"Ctrl+]":   "jump to the tag under cursor",
	}
}
