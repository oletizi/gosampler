package audio

type Sink interface {
	Write(b Buffer) error
	Close() error
}
