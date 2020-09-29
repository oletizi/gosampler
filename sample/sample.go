package sample

import "osampler"

type sample struct {
	filename string
}

func New(filename string) osampler.Sample {
	return sample{filename: filename}
}

func (s sample) Filename() string {
	return s.filename
}
