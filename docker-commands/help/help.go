package help

const (
	HelperName = "help"
)

type Helper struct{}

func NewHelper() *Helper {
	return &Helper{}
}

func (s *Helper) GetCommands() map[string]string {
	return map[string]string{
		"images":     "get all images commands",
		"hub":        "get all hub commands",
		"general":    "get all general commands",
		"containers": "get all container commands",
	}
}
