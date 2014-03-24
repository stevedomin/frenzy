package docker

import (
	"bytes"
	"fmt"
	"github.com/stevedomin/frenzy/pkg"
	"os/exec"
	"strings"
)

type Provider struct {
	Image string
}

func (p *Provider) Up(hostname string) (*pkg.NodeInfo, error) {
	fmt.Printf("[docker] up %s\n", hostname)
	var outBuf, errBuf bytes.Buffer

	err := p.cmd(
		"docker",
		[]string{
			"run",
			"-d",
			"-p", "22",
			"--name", hostname,
			p.Image,
			"/usr/sbin/sshd", "-D",
		},
		&outBuf,
		&errBuf,
	)
	if err != nil {
		return nil, fmt.Errorf("[docker] Error while running container %s: %s", hostname, errBuf.String())
	}

	nodeInfo := &pkg.NodeInfo{}
	nodeInfo.ID = strings.Trim(outBuf.String(), "\n")

	err = p.cmd(
		"docker",
		[]string{
			"port",
			nodeInfo.ID,
			"22",
		},
		&outBuf,
		&errBuf,
	)
	if err != nil {
		return nil, fmt.Errorf("[docker] Error while getting container port %s: %s", hostname, errBuf.String())
	}

	nodeInfo.Host = "127.0.0.1"
	nodeInfo.Port = strings.Split(strings.Trim(outBuf.String(), "\n"), ":")[1]

	return nodeInfo, nil
}

func (p *Provider) Stop(hostname, ID string) error {
	return nil
}

func (p *Provider) Destroy(hostname, ID string) error {
	fmt.Printf("[docker] destroying %s, id: %s\n", hostname, ID[:12])
	var outBuf, errBuf bytes.Buffer

	err := p.cmd(
		"docker",
		[]string{
			"kill",
			ID,
		},
		&outBuf,
		&errBuf,
	)
	if err != nil {
		return fmt.Errorf(strings.Trim(errBuf.String(), "\n"))
	}

	err = p.cmd(
		"docker",
		[]string{
			"rm",
			ID,
		},
		&outBuf,
		&errBuf,
	)
	if err != nil {
		return fmt.Errorf(strings.Trim(errBuf.String(), "\n"))
	}

	return nil
}

func (p *Provider) cmd(name string, args []string, outBuf, errBuf *bytes.Buffer) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
