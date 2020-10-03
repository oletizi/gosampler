package audio

type Buffer interface {
	Data() []float64
	Resize(size int)
	Size() int
}

type buffer struct {
	data []float64
}

func NewBuffer(size int) Buffer {
	return &buffer{
		make([]float64, size),
	}
}

func (b *buffer) Data() []float64 {
	return b.data
}

func (b *buffer) Size() int {
	return len(b.data)
}

func (b *buffer) Resize(size int) {
	if size < len(b.data) {
		b.data = b.data[0:size]
	}
	if size > len(b.data) {
		puffy := make([]float64, size-len(b.data))
		b.data = append(b.data, puffy...)
	}
}
