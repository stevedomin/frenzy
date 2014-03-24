package main

import (
	"encoding/json"
	"fmt"
	"github.com/stevedomin/cli"
	"github.com/stevedomin/frenzy/pkg"
	"github.com/stevedomin/frenzy/pkg/commands"
	"github.com/stevedomin/frenzy/pkg/environment"
	"io/ioutil"
	"os"
)

var privateKeyFile string = "./.frenzy/frenzy_insecure_key"

const privateKey string = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA5TDpI5k4q229isBtXVSNwtNO146lqQKqsW/ETSg1Il82PT25
lJO7AByBp0odDwIgWs4UoAbo4KYUkzIb7+94OfZhOLPW9MCT39hDzc3AsJuoODjG
s+55sXkfvMgEeFlAaZrmFsGjM0mp+pUpQT4PIt8k4f6nEZbeotYUTNSBZNoo02Gd
uW/UwpxvYmoBx119nGQOSIw0nv8QB8cYVr4hbYfQmhG2yryKqhmtDCM631uf7M7f
/o0DduwGy/OZ7CznStvHDfPAloQShd4MPanWxS7L+0z259hoJob8bLREK/boR1X0
GEDzL/L6xuPYqyl2/5VvAr5ezFfFMSPt/5iR6wIDAQABAoIBAQC6XYJGkzI0m8DT
0dNcYAQCF2d1+qWUf/mi3Ppnrzk9oORu+gEs4s/dsFBxYt+sM5Nxoz+8PMIi4om6
g7WZ7kT6BPFbdUlmri3QiH/iGtwEAB7S0MAq0dEc0hxPmumfnxu0g+NzP7PgYZcZ
jy8DwV33gjHwnuzlbaPOD3xkWSx6fy0aDiMnmhkPZb5Fy80/8cpYxefmCX5dz4ps
GawRDReBIN1fixiUAfsUsmIjpwYfc8YFmqt4XkvnTOpG4CwAHng+QP0fIKIuvCSR
MauJ/GRXwRVPCC4b9BSp9Xm/NLCIfsCVMvifQVYA0ibCyMAefpeR92vHJbVkw+WS
8mkrwysxAoGBAPuUzs8d29+/OCJIhAunpE5J3sv/cQ2EHcpJe95B4D7U+Oj9SHgr
tEisC0aq6R+GZK6UzonI+ujA+HXJjK6dQmQmjbIMg2J078n1hOht358EvRKsnwo8
Xj4o4d5bRb+uchFYUx/AT9kQnckoxYfl8Qhs1V1/YS/saYO+IMoreMLTAoGBAOk3
bdX+EIRLUMBRamSJNcogl+KIDG1YStg/N9e6xT0pQcR67m9mDTIX5Mt9gtP5mFMX
BYi8Tyb8hm4+xPbJuSk0xrL+LGvL1ta9nkIotXV8HhgOpGhPSobFfhe6hN71qmtg
YhhBAjPOwL0hWyuxgcks3vyqqpz4llbmJ6mr7xWJAoGBAMcyjQl7V+PycQzcJAli
ZHtEjC80A5yzFi9cPcK+oEK/uJIqMh5MZIQCDS+YFdvLOp7s3hhE1T5DxLbmrgh4
JeBMknb+52ymsFJVnzW2AZDUXKyTl52wLOLE1gqMdE6QXmsTZ0XFrLNvL6/eI4E1
9MI6Ajr0p8wdQXJ5sVbCUuzXAoGAVvP+tMG0gM7f/cSdSXzLHGDIoIKfN4mq5jOz
63BMqAAg8FPxYIGJO/siqChUxhoSjDBSNOgZDu0qZiJrJ6SQr+doikYwY8SthoVO
Y1ZwSayMP+X9sJKXs+YgnFJDximMb9qQ2IAshQp6XEn6hzD4tfxP657qW/pne2D2
utFLJBkCgYEAlH+tHtEtfQCXnuRu+3+z0ghNDPk1dwIx7OVfo/BePStvpeL32uMS
EoIKY3/tT6SaEvdFQOojYHoLJCdPqsoYVlX0eeanYbFCVdGo/gR1HHnowzSR+lQP
evYRSe4TuyhP8954HQXgSbAiNfMbNhW0CzSsguBWsoKmwbInsinwLTU=
-----END RSA PRIVATE KEY-----
`

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := createFrenzyDirStructure(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	env := environment.NewEnvironment()
	env.Bootstrap(config)
	env.LoadState()

	rootCmd := cli.NewCommand("frenzy")
	rootCmd.HelpFunc = Usage()

	rootCmd.AddCommands(
		commands.Destroy(env),
		commands.Provision(env),
		commands.Stop(env),
		commands.SSH(env),
		commands.Status(env),
		commands.Up(env),
		commands.Version(),
	)
	rootCmd.Execute(nil)
}

func loadConfig() (*pkg.Config, error) {
	configfile, err := ioutil.ReadFile("./Frenzyfile")
	if err != nil {
		return nil, fmt.Errorf("Error while reading Frenzyfile: %s", err)
	}

	var c pkg.Config
	err = json.Unmarshal(configfile, &c)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling Frenzyfile: %s", err)
	}

	return &c, nil
}

func createFrenzyDirStructure() error {
	if _, err := os.Stat("./.frenzy"); os.IsNotExist(err) {
		err := os.Mkdir("./.frenzy", 0777)
		if err != nil {
			fmt.Errorf("Error while creating .frenzy folder: %s", err)
		}
	}

	if _, err := os.Stat(privateKeyFile); os.IsNotExist(err) {
		file, err := os.Create(privateKeyFile)
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString(privateKey)
	}

	return nil
}

func Usage() func() {
	return func() {
		fmt.Println(`Usage: frenzy [<options>] <command> [<args>]

Commands:
    destroy                 Destroy all nodes.
    provision               Provision all nodes.
    ssh                     Login to node via SSH.
    status                  Display current state of nodes.
    up                      Start nodes.
    version                 Display version information.
`)
	}

}
