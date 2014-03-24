package commands

import (
	"github.com/stevedomin/cli"
	"github.com/stevedomin/frenzy/pkg"
	"github.com/stevedomin/frenzy/pkg/environment"
	"log"
	"sync"
)

var (
	flagProvision bool
)

func Up(env *environment.Environment) *cli.Command {
	upCmd := cli.NewCommand("up")
	upCmd.Flags.BoolVar(&flagProvision, "provision", true, "Provision")
	upCmd.HandlerFunc = func(args []string) {
		var wg sync.WaitGroup
		for _, node := range env.Nodes {
			wg.Add(1)
			go func(node *pkg.Node) {
				defer wg.Done()

				if node.Status != "running" {
					nodeInfo, err := node.Provider.Up(node.Hostname)
					if err != nil {
						log.Println(err)
						return
					}

					node.ID = nodeInfo.ID
					node.Host = nodeInfo.Host
					node.Port = nodeInfo.Port
					node.Status = "running"
				} else {
					log.Printf("[%s] already running", node.Hostname)
				}

				if flagProvision {
					for _, provisioner := range node.Provisioners {
						provisioner.Run(node)
					}
				}
			}(node)
		}
		wg.Wait()

		if err := env.SaveState(); err != nil {
			log.Println(err)
		}
	}
	return upCmd
}
