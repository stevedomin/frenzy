package commands

import (
	"fmt"
	"github.com/stevedomin/cli"
	"github.com/stevedomin/frenzy/pkg"
)

func Version() *cli.Command {
	versionCmd := cli.NewCommand("version")
	versionCmd.HandlerFunc = func(args []string) {
		fmt.Printf("frenzy v%s\n", pkg.Version)
	}
	return versionCmd
}
