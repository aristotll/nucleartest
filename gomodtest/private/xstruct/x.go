package xstruct

type private struct {
	x string
	Y string
}

func Retpriv() *private {
	return &private{"123", "456"}
}
