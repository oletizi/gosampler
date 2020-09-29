package osampler

type Note interface {
	Value() int
	Name() string
}
type Instrument interface {
	NoteOn(note *Note)
	NoteOff(note *Note)
}

type Player interface {
	Play()
	Stop()
}

type Sample interface {
	Filename() string
}

type Config interface {
	SamplesFor(note Note) []Sample
}
