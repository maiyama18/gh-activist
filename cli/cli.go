package cli

const (
	Success = iota
)

type Cli struct {}

func New() (*Cli, error) {
	return &Cli{}, nil
}

func (c *Cli) Run() int {
	return Success
}
