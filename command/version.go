package command

import (
	"fmt"
	log "github.com/Sirupsen/logrus"

	"github.com/johnbellone/skel/skel"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Skel",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("version called")
		fmt.Printf("skel %s\r\n", skel.SemVersion)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
