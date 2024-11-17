package register

const (
	RegisterName = "register"
)

type Register struct{}

func NewRegister() *Register {
	return &Register{}
}

func (s *Register) GetCommands() map[string]string {
	return map[string]string{
		":reg[isters]": "show registers content",
		`"xy`:          "yank into register x",
		`"xp`:          "paste contents of register x",
		`"+y`:          "yank into the system clipboard register",
		`"+p`:          "paste from the system clipboard register",
	}
}
