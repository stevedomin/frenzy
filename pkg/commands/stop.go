package commands

import (
	"fmt"
	"github.com/stevedomin/cli"
	"github.com/stevedomin/frenzy/pkg/environment"
)

func Stop(env *environment.Environment) *cli.Command {
	stopCmd := cli.NewCommand("stop")
	stopCmd.HandlerFunc = func(args []string) {
		fmt.Printf("not yet implemented\n")
		// Docker commit container?
		// Docker stop
		// defer env.SaveState()
	}
	return stopCmd
}
