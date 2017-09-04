package glPget

type Options struct {
	Help    bool `short:"h" long:"help" description:"show this message"`
	Version bool `short:"v" long:"version" description:"show the version"`

	Trace bool `long:"trace" description:"show stack talace"`
	Procs bool `short:"p" long:"procs" description:"set pararell downlods procs"`
}
