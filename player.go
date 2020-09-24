package osampler

import (
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"

	"osampler/audio"
	"osampler/util"
)

func PlayFile(out audio.Out, path string) error {
	streamer, _, err := streamerFor(path)
	if err == nil {
		out.Play(streamer)
		streamer.Close()
	}
	return err
}

func streamerFor(path string) (beep.StreamSeekCloser, beep.Format, error) {
	f, err := os.Open(path)
	util.HandleFatal(err)
	return mp3.Decode(f)
}
