package midi

type Context interface {
	Tempo() float64
	PPQ() float64
	MsPerTick() float64
	MillisToTicks() float64
}

type context struct {
	tempo float64
	ppq   float64
}

func (c *context) MsPerTick() float64 {
	return 60 * 1000 / (c.tempo * c.ppq)
}

func (c *context) MillisToTicks() float64 {
	panic("implement me")
}

func (c *context) Tempo() float64 {
	return c.tempo
}

func (c *context) PPQ() float64 {
	return c.ppq
}

func NewContext(tempo float64, ppq float64) Context {
	return &context{tempo, ppq}
}
