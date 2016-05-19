package command

type SyncCommand struct {
}

func (c *SyncCommand) Run(args []string) int {
	return 0
}

func (c *SyncCommand) Help() string {
	return ""
}

func (c *SyncCommand) Synopsis() string {
	return ""
}
