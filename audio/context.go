package audio

type Context interface {
	SampleRate() int
	BitDepth() BitDepth
	ChannelCount() int
}

type context struct {
	sampleRate   int
	bitDepth     BitDepth
	channelCount int
}

// XXX: Find a decent naming convention for convenience context functions
func NewCDContext() Context {
	return NewContext(44100, NewBitDepth16(), 2)
}

func NewContext(sampleRate int, bitDepth BitDepth, channelCount int) Context {
	return context{sampleRate, bitDepth, channelCount}
}

func (c context) SampleRate() int {
	return c.sampleRate
}

func (c context) BitDepth() BitDepth {
	return c.bitDepth
}

func (c context) ChannelCount() int {
	return c.channelCount
}
