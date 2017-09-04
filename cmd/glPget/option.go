package glPget

import (
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Help    bool `short:"h" long:"help" description:"show this message"`
	Version bool `short:"v" long:"version" description:"show the version"`

	Trace bool `long:"trace" description:"show stack talace"`
	Procs bool `short:"p" long:"procs" description:"set pararell downlods procs"`
}

func (opts *Options) parse(argv []string) ([]string, error) {
	parser := flags.NewParser(opts, flags.None)
	args, err := parser.ParseArgs(argv)
	return args, nil
}
