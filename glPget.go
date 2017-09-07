package glPget

import (
	"fmt"
	"os"
	"runtime"

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
	fmt.Println(glp.Procs)
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

	if len(o) == 0 {
		return glp.showHelp(opts)
	}

	if opts.Help {
		return glp.showHelp(opts)
	}

	// set default procs
	if opts.Procs == 0 {
		opts.Procs = runtime.NumCPU()
	}

	glp.args = o
	return nil
}

func (glp glPget) showHelp(opts *Options) ignoreError {
	os.Stdout.Write(opts.usage())
	return glp.makeIgnoreError()
}
