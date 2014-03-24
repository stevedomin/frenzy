package pkg

type Provisioner interface {
	Run(node *Node)
}
