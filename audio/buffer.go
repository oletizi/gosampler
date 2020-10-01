package audio

type Buffer interface {
	Context() Context
	Data() []float64
	Size() int
}

type buffer struct {
	context Context
	data    []float64
}

func (b buffer) Size() int {
	return len(b.data)
}

func NewBuffer(c Context, size int) Buffer {
	return buffer{
		c,
		make([]float64, size),
	}
}

func (b buffer) Context() Context {
	return b.context
}

func (b buffer) Data() []float64 {
	return b.data
}
