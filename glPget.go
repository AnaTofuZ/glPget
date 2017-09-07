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

// 純粋なerrorを取ってくるインターフェイス
// http://deeeet.com/writing/2016/04/25/go-pkg-errors/ を参照

type causer interface {
	Cause() error
}

type ignoreError struct {
	Msg []byte
}

func New() *glPget {

	return &glPget{}
}

func (glp *glPget) Run() error {
	err := glp.prepare(os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "stoping run")
	}
	return nil
}

// ignoreがErrorを持つように定義する
func (i *ignoreError) Error() string {
	return fmt.Sprintf("%s", i.Msg)
}

func (glp *glPget) prepare(argv []string) error {

	if err := glp.parseOptions(&glp.Options, argv); err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

func (glp *glPget) parseOptions(opts *Options, argv []string) error {
	o, err := opts.parse(argv)

	if err != nil {
		return exit.MakeDataErr(err)
	}

	if opts.Version {
		return glp.showVersion()
	}

	if len(o) == 0 || opts.Help {
		return glp.showHelp()
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

func (glp *glPget) showHelp() error {
	return &ignoreError{Msg: glp.Options.usage()}
}

func (glp *glPget) showVersion() error {
	return &ignoreError{Msg: []byte("glpget version" + version + "\n")}
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
