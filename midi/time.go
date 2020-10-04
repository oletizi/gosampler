package midi

type Clock interface {
	ElapsedMilliseconds() uint64
	ElapsedTicks() uint64
}
