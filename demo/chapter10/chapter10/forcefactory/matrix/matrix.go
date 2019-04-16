package matrix

type matrix struct {
	x int
}

func NewMatrix(x int) *matrix {
	return &matrix{x}
}
