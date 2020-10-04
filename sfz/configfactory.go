package sfz

import (
	"log"
	"path/filepath"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/oletizi/sfz-parser/go/parser"

	"osampler/instrument"
	"osampler/midi"
	"osampler/sample"
)

type samplerConfig struct {
	regions []*region
}

func (s *samplerConfig) SamplesFor(note midi.Note) []sample.Sample {
	var samples []sample.Sample
	for i := 0; i < len(s.regions); i++ {
		region := s.regions[i]
		// XXX: THer MUST be a more elegant way to do this
		lokey := region.lokey
		key := region.key
		hikey := region.hikey
		theSample := region.sample
		if (key != nil && note == key ||
			(lokey != nil && note.Value() >= lokey.Value() &&
				(hikey != nil && note.Value() <= hikey.Value()))) &&
			theSample != nil {
			samples = append(samples, theSample)
		}
	}
	return samples
}

type region struct {
	sample         sample.Sample
	hikey          midi.Note
	key            midi.Note
	lokey          midi.Note
	hivel          int
	lovel          int
	pitchKeycenter int
}

//
// SFZ file format parsing
//

type sfzListener struct {
	*parser.BaseSfzListener

	cfg           *samplerConfig
	filepath      string
	basedir       string
	currentRegion *region
	currentOpcode string
}

func New(sfzFile string) (instrument.Config, error) {
	in, err := antlr.NewFileStream(sfzFile)
	if err != nil {
		return nil, err
	}
	lexer := parser.NewSfzLexer(in)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSfzParser(stream)
	cfg := &samplerConfig{}
	listener := &sfzListener{cfg: cfg, filepath: sfzFile, basedir: filepath.Dir(sfzFile)}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Sfz())

	return cfg, nil
}

func (s *sfzListener) ExitHeader(ctx *parser.HeaderContext) {
	header := ctx.GetText()
	switch header {
	case "region":
		region := &region{}
		s.currentRegion = region
		s.cfg.regions = append(s.cfg.regions, region)
	}
}

func (s *sfzListener) ExitOpcode(ctx *parser.OpcodeContext) {
	s.currentOpcode = ctx.GetText()
}

func (s *sfzListener) ExitValue(ctx *parser.ValueContext) {

	switch s.currentOpcode {

	case "sample":
		filename := filepath.Clean(filepath.Join(s.basedir, ctx.GetText()))
		theSample := sample.New(filename)
		s.currentRegion.sample = theSample

	case "key":
		key, err := resolveNote(ctx.GetText())
		if err != nil {
			log.Printf("Error parsing note value: %v", err)
		}
		s.currentRegion.key = key
	case "lokey":
		lokey, err := resolveNote(ctx.GetText())
		if err != nil {
			log.Printf("Error parsing note value: %v", err)
		}
		s.currentRegion.lokey = lokey

	case "hikey":
		hikey, err := resolveNote(ctx.GetText())
		if err != nil {
			log.Printf("Error parsing note value: %v", err)
		}
		s.currentRegion.hikey = hikey

	}
}

func resolveNote(svalue string) (midi.Note, error) {
	ivalue, err := strconv.Atoi(svalue)
	if err != nil {
		return nil, err
	}
	return midi.NewNote(ivalue)
}
