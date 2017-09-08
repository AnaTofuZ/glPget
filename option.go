package glPget

import (
	"bytes"
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

type Options struct {
	Help    bool `short:"h" long:"help" description:"show this message"`
	Version bool `short:"v" long:"version" description:"show the version"`

	Trace bool `long:"trace" description:"show stack trace"`
	Procs int  `short:"p" long:"procs" description:"set pararell downlods procs"`
}

func (opts *Options) parse(argv []string) ([]string, error) {
	parser := flags.NewParser(opts, flags.None)
	args, err := parser.ParseArgs(argv)
	if err != nil {
		os.Stderr.Write(opts.usage())
		return nil, errors.Wrap(err, "invalid command line")
	}
	return args, nil
}

func (opts *Options) usage() []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf,
		`glPget`+version+`
Usage: glPget [options] URL

Options:
-h, --help 			show this message
-v, --version		show the  version
-p,	--procs			set pararell downlods procs	
--trace				show stack trace
`)

	return buf.Bytes()
}
