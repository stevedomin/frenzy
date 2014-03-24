package pkg

import (
	"github.com/stevedomin/frenzy/pkg/communicators/ssh"
)

type NodeInfo struct {
	ID       string
	Hostname string
	Host     string
	Port     string
}

type Node struct {
	Communicator Communicator
	ID           string
	Status       string
	Hostname     string
	Host         string
	Port         string
	Provider     Provider
	Provisioners []Provisioner
}

// TODO Change this
func (n *Node) GetCommunicator() Communicator {
	if n.Communicator == nil {
		// Private key path shouldn't be hardcoded that way...
		n.Communicator = ssh.NewSSH(n.Host, n.Port, "./.frenzy/frenzy_insecure_key")
	}
	return n.Communicator
}

func (n *Node) TruncID() string {
	if n.ID != "" {
		return n.ID[:12]
	} else {
		return n.ID
	}
}
