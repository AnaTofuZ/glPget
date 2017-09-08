package glPget

import (
	"fmt"
	"os"
	"runtime"

	"github.com/asaskevich/govalidator"

	"github.com/Code-Hex/exit"
	"github.com/pkg/errors"
)

const (
	version = "0.01"
)

type glPget struct {
	Options

	URL string
}

func New() *glPget {

	return &glPget{}
}

func (glp *glPget) Run() int {
	if e := glp.run(); e != nil {
		exitCode, err := glp.ErrTrap(e)
		if glp.Trace {
			fmt.Fprintf(os.Stderr, "Error:%+v\n", e)
		} else {
			fmt.Fprintf(os.Stderr, "Error:%v\n", err)
		}
		return exitCode
	}
	return 0
}

func (glp *glPget) run() error {
	err := glp.prepare(os.Args[1:])
	if err != nil {
		return err
	}
	return nil
}

func (glp *glPget) prepare(argv []string) error {

	if err := glp.parseOptions(&glp.Options, argv); err != nil {
		return errors.Wrap(err, "faild to parse")
	}

	return nil
}

func (glp *glPget) parseOptions(opts *Options, argv []string) error {
	o, err := opts.parse(argv)

	if err != nil {
		return exit.MakeDataErr(err)
	}

	if opts.Version {
		os.Stdout.Write([]byte("glpget version" + version + "\n"))
		os.Exit(0)
	}

	if opts.Help {
		os.Stdout.Write(opts.usage())
		os.Exit(0)
	}

	if len(o) == 0 {
		return exit.MakeUsage(errors.New(string(opts.usage())))
	}

	// set default procs
	if opts.Procs == 0 {
		opts.Procs = runtime.NumCPU()
	}

	if err := glp.setURL(o); err != nil {
		return errors.Wrap(err, "url is not found")
	}

	return nil
}

func (glp *glPget) setURL(args []string) error {

	for _, argv := range args {
		if govalidator.IsURL(argv) {
			glp.URL = argv
			break
		}
	}

	if len(glp.URL) < 1 {
		return errors.New("urls not found")
	}

	return nil
}
