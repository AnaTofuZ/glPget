package glPget

const (
	version = "0.01"
)

type glPget struct {
	Options

	URL []string
}

func New() *glPget {
	return &glPget{}
}

func parseOptions(opts *Options, argv []string) ([]string, error) {
	o, err := opts.parse(argv)

	return o, nil
}
