package cli

type Cli struct {
	Name     string
	HelpName string
	Version  string
	Usage    string
	Commands []Command
	Options  []Option
}

func New(name string) *Cli {
	if name == "" {
		panic("can't construct an app without a name")
	}
	c := &Cli{Name: name}
	c.setUp()
	return c
}

func (c *Cli) setUp() {
	help := Option{
		Name:  "help",
		Short: "h",
		Usage: "Display this help message",
	}

	version := Option{
		Name:  "version",
		Short: "v",
		Usage: "Display this application version",
	}

	c.Options = append(c.Options, help)
	c.Options = append(c.Options, version)
}

func (c *Cli) Run() {
	c.printHelp()
}
