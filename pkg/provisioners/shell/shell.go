package shell

import (
	"fmt"
	"github.com/stevedomin/frenzy/pkg"
)

type Provisioner struct {
	Inline []string
}

func (p *Provisioner) Run(node *pkg.Node) {
	fmt.Printf("[%s] Running inline SSH provisioner\n", node.Hostname)

	for _, inline := range p.Inline {
		output := node.GetCommunicator().Run(inline)
		fmt.Printf("[%s] %s", node.Hostname, output)
	}
}
