package shell

import (
	"github.com/stevedomin/frenzy/pkg"
	"log"
)

type Provisioner struct {
	Inline []string
}

func (p *Provisioner) Run(node *pkg.Node) {
	log.Printf("[%s] Running inline SSH provisioner", node.Hostname)

	for _, inline := range p.Inline {
		output := node.GetCommunicator().Run(inline)
		log.Printf("[%s] %s", node.Hostname, output)
	}
}
