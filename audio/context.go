package audio

type Context interface {
	SampleRate() int
	BitDepth() int
	ChannelCount() int
}

type context struct {
	sampleRate   int
	bitDepth     int
	channelCount int
}

func NewContext(sampleRate int, bitDepth int, channelCount int) Context {
	return context{sampleRate, bitDepth, channelCount}
}

func (c context) SampleRate() int {
	return c.sampleRate
}

func (c context) BitDepth() int {
	return c.bitDepth
}

func (c context) ChannelCount() int {
	return c.channelCount
}
