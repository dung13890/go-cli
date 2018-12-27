package cli

type Command struct {
	Name     string
	HelpName string
	Usage    string
	Options  []Option
}

type Option struct {
	Name     string
	Short    string
	Usage    string
	Help     string
	Variable bool
}
