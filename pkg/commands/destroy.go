package commands

import (
	"github.com/stevedomin/cli"
	"github.com/stevedomin/frenzy/pkg"
	"github.com/stevedomin/frenzy/pkg/environment"
	"log"
	"sync"
)

func Destroy(env *environment.Environment) *cli.Command {
	destroyCmd := cli.NewCommand("destroy")
	destroyCmd.HandlerFunc = func(args []string) {
		var wg sync.WaitGroup
		for _, node := range env.Nodes {
			wg.Add(1)
			go func(node *pkg.Node) {
				defer wg.Done()
				if node.Status != "not running" {
					err := node.Provider.Destroy(node.Hostname, node.ID)
					if err != nil {
						log.Println(err)
					}
				} else {
					log.Printf("[%s] already destroyed", node.Hostname)
				}

				err := env.DestroyState()
				if err != nil {
					log.Println(err)
				}
			}(node)
		}
		wg.Wait()
	}

	return destroyCmd
}
