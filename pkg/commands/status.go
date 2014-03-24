package commands

import (
	"fmt"
	"github.com/stevedomin/cli"
	"github.com/stevedomin/frenzy/pkg/environment"
	"github.com/stevedomin/termtable"
)

func Status(env *environment.Environment) *cli.Command {
	statusCmd := cli.NewCommand("status")
	statusCmd.HandlerFunc = func(args []string) {
		t := termtable.NewTable(nil, &termtable.TableOptions{
			Padding:      2,
			UseSeparator: false,
		})
		t.SetHeader([]string{"HOSTNAME", "STATUS", "CONTAINER ID", "PORT"})

		for _, node := range env.Nodes {
			t.AddRow([]string{
				node.Hostname,
				node.Status,
				node.TruncID(),
				node.Port,
			})
		}

		fmt.Println(t.Render())
	}
	return statusCmd
}
