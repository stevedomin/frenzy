package commands

import (
	"github.com/stevedomin/cli"
	"github.com/stevedomin/frenzy/pkg/environment"
	"log"
)

func Stop(env *environment.Environment) *cli.Command {
	stopCmd := cli.NewCommand("stop")
	stopCmd.HandlerFunc = func(args []string) {
		log.Printf("not yet implemented")
		// Docker commit container?
		// Docker stop
		// defer env.SaveState()
	}
	return stopCmd
}
