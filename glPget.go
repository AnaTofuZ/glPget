package glPget

import (
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

	args []string
	URL  string
}

// 純粋なerrorを取ってくるインターフェイス
// http://deeeet.com/writing/2016/04/25/go-pkg-errors/ を参照

type causer interface {
	Cause() error
}

type ignoreError struct {
	err error
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

// usage時などにエラー表示させないようにする関数
func (glp glPget) ErrTop(err error) error {
	for e := err; e != nil; {
		switch e.(type) {
		case ignoreError:
			return nil
		case causer:
			e = e.(causer).Cause()
		default:
			return e
		}
	}
	return nil
}

// ignoreがErrorを持つように定義する
func (i ignoreError) Error() string {
	return i.err.Error()
}

func (glp glPget) makeIgnoreError() ignoreError {
	return ignoreError{
		err: errors.New("this is ignore error message"),
	}
}

func (glp *glPget) prepare(argv []string) error {

	if err := glp.parseOptions(&glp.Options, argv); err != nil {
		return glp.ErrTop(err)
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

	if err := glp.setURL(); err != nil {
		return errors.Wrap(err, "url is not found")
	}

	glp.args = o
	return nil
}

func (glp glPget) showHelp() ignoreError {
	os.Stdout.Write(glp.Options.usage())
	return glp.makeIgnoreError()
}

func (glp glPget) showVersion() ignoreError {
	os.Stdout.Write([]byte("glpget version" + version + "\n"))
	return glp.makeIgnoreError()
}

func (glp *glPget) setURL() error {

	for _, argv := range glp.args {
		if govalidator.IsURL(argv) {
			glp.URL = argv
		}
	}

	if len(glp.URL) < 1 {
		return errors.New("urls not found")
	}

	return nil
}
