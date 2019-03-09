package cli

type ExitCode int
const (
	Success ExitCode = iota
)

type Cli struct {}

func New() (*Cli, error) {
	return &Cli{}, nil
}

func (c *Cli) Run() ExitCode {
	return Success
}
