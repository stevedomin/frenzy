package environment

import (
	"encoding/json"
	"fmt"
	"github.com/stevedomin/frenzy/pkg"
	"github.com/stevedomin/frenzy/pkg/providers/docker"
	"github.com/stevedomin/frenzy/pkg/provisioners/shell"
	"io/ioutil"
	"log"
	"os"
)

const frenzyStatefile = "./.frenzy/state"

type Environment struct {
	Nodes []*pkg.Node
	State *State
}

func NewEnvironment() *Environment {
	e := &Environment{
		Nodes: []*pkg.Node{},
		State: &State{},
	}
	return e
}

func (e *Environment) Bootstrap(config *pkg.Config) {
	var defaultProvider pkg.Provider
	if config.Defaults.Provider.Docker != nil {
		d := config.Defaults.Provider.Docker
		defaultProvider = &docker.Provider{
			Image: d.Image,
		}
	}

	var defaultProvisioners []pkg.Provisioner
	for _, p := range config.Defaults.Provisioners {
		var provisioner pkg.Provisioner

		switch p.Type {
		case "shell":
			provisioner = &shell.Provisioner{
				Inline: p.Inline,
			}
		}

		defaultProvisioners = append(defaultProvisioners, provisioner)
	}

	for _, n := range config.Nodes {
		node := &pkg.Node{
			Hostname: n.Hostname,
			Status:   "not running",
		}

		if n.Provider == nil {
			node.Provider = defaultProvider
		} else {
			// load provider
		}

		if len(n.Provisioners) <= 0 {
			node.Provisioners = defaultProvisioners
		} else {
			// load provisioners
		}

		// log.Printf("[%s] Node: %+v", node.Hostname, node)
		e.Nodes = append(e.Nodes, node)
	}
}

type State struct {
	Nodes []*NodeState `json:"nodes"`
}

type NodeState struct {
	Hostname string `json:"hostname"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Status   string `json:"status"`
	ID       string `json:"id"`
}

func (e *Environment) LoadState() {
	if _, err := os.Stat(frenzyStatefile); os.IsNotExist(err) {
		return
	}

	b, err := ioutil.ReadFile(frenzyStatefile)
	if err != nil {
		log.Fatalf("Error while loading state: %s", err)
	}

	err = json.Unmarshal(b, e.State)
	if err != nil {
		log.Fatalf("Error while unmarshalling state: %s", err)
	}
	for _, nodeState := range e.State.Nodes {
		for _, node := range e.Nodes {
			if nodeState.Hostname == node.Hostname {
				node.ID = nodeState.ID
				node.Host = nodeState.Host
				node.Port = nodeState.Port
				node.Status = nodeState.Status
			}
		}
	}
}

func (e *Environment) SaveState() error {
	e.State.Nodes = []*NodeState{}
	for _, node := range e.Nodes {
		n := &NodeState{
			ID:       node.ID,
			Hostname: node.Hostname,
			Host:     node.Host,
			Port:     node.Port,
			Status:   node.Status,
		}
		e.State.Nodes = append(e.State.Nodes, n)
	}

	b, err := json.Marshal(*e.State)
	if err != nil {
		return fmt.Errorf("Error while marshalling state: %s", err)
	}
	err = ioutil.WriteFile(frenzyStatefile, b, 0644)
	if err != nil {
		return fmt.Errorf("Error while writing state: %s", err)
	}
	return nil
}

func (e *Environment) DestroyState() error {
	if _, err := os.Stat(frenzyStatefile); os.IsNotExist(err) {
		return nil
	}

	err := os.Remove(frenzyStatefile)
	if err != nil {
		return err
	}
	return nil
}
