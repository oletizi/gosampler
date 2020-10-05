package sfz

import (
	"log"
	"path/filepath"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/oletizi/sfz-parser/go/parser"

	"osampler/instrument"
	"osampler/midi"
)

type samplerConfig struct {
	regions []*region
}

func (s *samplerConfig) FilesFor(note midi.Note) []string {
	var files []string
	for i := 0; i < len(s.regions); i++ {
		region := s.regions[i]
		// XXX: THer MUST be a more elegant way to do this
		lokey := region.lokey
		key := region.key
		hikey := region.hikey
		filepath := region.filepath
		if (key != nil && note == key ||
			(lokey != nil && note.Value() >= lokey.Value() &&
				(hikey != nil && note.Value() <= hikey.Value()))) &&
			filepath != "" {
			files = append(files, filepath)
		}
	}
	return files
}

type region struct {
	filepath       string
	hikey          midi.Note
	key            midi.Note
	lokey          midi.Note
	hivel          int
	lovel          int
	pitchKeycenter int
}

//
// SFZ filepath format parsing
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
		path := filepath.Clean(filepath.Join(s.basedir, ctx.GetText()))
		s.currentRegion.filepath = path

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
