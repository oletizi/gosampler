package audio

import "github.com/faiface/beep"

type Out interface {
	Play(beep.StreamSeekCloser)
}
