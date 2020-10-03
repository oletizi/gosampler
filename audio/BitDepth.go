package audio

import "math"

type BitDepth interface {
	Depth() int
	MaxInt() float64
}

type bitDepth struct {
	depth  int
	maxInt float64
}

func NewBitDepth16() BitDepth {
	return &bitDepth{16, float64(math.MaxInt16)}
}

func (b *bitDepth) Depth() int {
	return b.depth
}

func (b *bitDepth) MaxInt() float64 {
	return b.maxInt
}
