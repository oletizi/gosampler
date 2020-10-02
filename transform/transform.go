package transform

type Transform interface {
	CalculateBuffer() int
	Close() error
}

type BaseTransform struct{}

func (n BaseTransform) CalculateBuffer() int { return 0 }
func (n BaseTransform) Close() error         { return nil }
