package midi

type Context interface {
	Tempo() float64
	PPQ() float64
	MsPerTick() float64
}

type BaseContext struct {
	tempo float64
	ppq   float64
}

func (c *BaseContext) MsPerTick() float64 {
	return 60 * 1000 / (c.tempo * c.ppq)
}

func (c *BaseContext) Tempo() float64 {
	return c.tempo
}

func (c *BaseContext) PPQ() float64 {
	return c.ppq
}
