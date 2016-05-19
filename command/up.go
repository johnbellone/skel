package command

type UpCommand struct {
}

func (c *UpCommand) Run(args []string) int {
	return 0
}

func (c *UpCommand) Help() string {
	return ""
}

func (c *UpCommand) Synopsis() string {
	return ""
}
