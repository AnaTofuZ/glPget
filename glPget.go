package glPget

import (
	"fmt"
	"os"

	"github.com/Code-Hex/exit"
	"github.com/pkg/errors"
)

const (
	version = "0.01"
)

type glPget struct {
	Options

	args []string
	URL  []string
}

func New() *glPget {

	fmt.Println("hoge")
	return &glPget{}
}

func (glp *glPget) Run() {
	//	arg, err := glp.parseOptions(&glp.Options, os.Args[1:])
	err := glp.prepare(os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "stoping run")
	}
	fmt.Println(glp.args)
}

func (glp *glPget) prepare(argv []string) error {

	if err := glp.parseOptions(&glp.Options, argv); err != nil {
		return errors.Wrap(err, "faild to parse command line args")
	}

	if glp.Help {
		//		return glp.makeIgnore
	}

	return nil
}

func (glp *glPget) parseOptions(opts *Options, argv []string) error {
	o, err := opts.parse(argv)

	if err != nil {
		return exit.MakeDataErr(err)
	}

	glp.args = o
	return nil
}
