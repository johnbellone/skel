package command

type NewCommand struct {
}

func (c *NewCommand) Run(args []string) int {
	return 0
}

func (c *NewCommand) Help() string {
	return ""
}

func (c *NewCommand) Synopsis() string {
	return ""
}
