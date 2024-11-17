package general

const (
	GeneralName = "general"
)

type General struct{}

func NewGeneral() *General {
	return &General{}
}

func (s *General) GetCommands() map[string]string {
	return map[string]string{
		"docker -d":     "start docker daemon",
		"docker --help": "get help with docker",
		"docker info":   "display system-wide information",
	}
}
