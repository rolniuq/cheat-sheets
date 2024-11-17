package searchmul

const (
	SearchMulName = "search_mul"
)

type SearchMul struct{}

func NewSearchMul() *SearchMul {
	return &SearchMul{}
}

func (s *SearchMul) GetCommands() map[string]string {
	return map[string]string{
		":vim[grep] /pattern/ {`{file}`}": "search for pattern in multiple files",
		":cn[ext]":                        "jump to the next match",
		":cp[revious]":                    "jump to the previous match",
		":cope[n]":                        "open a window containing the list of matches",
		":ccl[ose]":                       "close the quickfix window",
	}
}
