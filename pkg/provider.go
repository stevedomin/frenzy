package pkg

type Provider interface {
	Up(hostname string) (*NodeInfo, error)
	Stop(hostname, ID string) error
	Destroy(hostname, ID string) error
}
