package sample

type Sample interface{}

type s struct{}

func New(filepath string) Sample {
	return &s{}
}
