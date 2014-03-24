package pkg

type Config struct {
	Defaults *cfgDefaults
	Nodes    []*cfgNode
}

type cfgDefaults struct {
	Provider     *cfgProvider
	Provisioners []*cfgProvisioner
}

type cfgNode struct {
	Hostname     string
	Provider     *cfgProvider
	Provisioners []*cfgProvisioner
}

type cfgProvider struct {
	Docker *cfgDockerProvider
}

type cfgDockerProvider struct {
	Image string
}

type cfgProvisioner struct {
	Type   string
	Inline []string
}
