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
