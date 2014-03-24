package commands

import (
	"fmt"
	"github.com/stevedomin/cli"
	"github.com/stevedomin/frenzy/pkg"
	"github.com/stevedomin/frenzy/pkg/environment"
	"os"
	"os/exec"
)

func SSH(env *environment.Environment) *cli.Command {
	sshCmd := cli.NewCommand("ssh")
	sshCmd.HandlerFunc = func(args []string) {
		if len(args) == 0 {
			fmt.Println("You need to specify the node you want to SSH into")
			fmt.Println("Example:")
			fmt.Println("	$ frenzy ssh node01")
			fmt.Println("")
			return
		}

		hostname := args[0]
		var node *pkg.Node

		for _, n := range env.Nodes {
			if n.Hostname == hostname {
				node = n
			}
		}

		if node.Status != "running" {
			fmt.Printf("[%s] Can't SSH. The node is not running.\n", node.Hostname)
			return
		}

		cmd := exec.Command(
			"ssh",
			"-o", "StrictHostKeyChecking=no",
			"-o", "UserKnownHostsFile=/dev/null",
			"-i", "keys/frenzy_insecure_key",
			fmt.Sprintf("frenzy@%s", node.Host),
			"-p", node.Port,
		)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
	return sshCmd
}
