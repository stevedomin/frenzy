package ssh

import (
	"bytes"
	"code.google.com/p/go.crypto/ssh"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type SSH struct {
	Connected bool
	Config    *SSHConfig
	Client    *ssh.ClientConn
}

type SSHConfig struct {
	User           string
	Host           string
	Port           string
	PrivateKeyPath string
}

func NewSSH(host, port, privateKeyPath string) *SSH {
	s := &SSH{
		Connected: false,
		Config: &SSHConfig{
			User:           "frenzy",
			Host:           host,
			Port:           port,
			PrivateKeyPath: privateKeyPath,
		},
	}
	return s
}

func (s *SSH) Run(command string) string {
	if !s.Connected {
		s.Connect()
	}

	session, err := s.Client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		panic("Failed to run: " + err.Error())
	}
	return b.String()
}

func (s *SSH) Connect() {
	block, _ := pem.Decode(loadPrivateKey(s.Config.PrivateKeyPath))
	rsakey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	clientKey := &keychain{rsakey}

	config := &ssh.ClientConfig{
		User: s.Config.User,
		Auth: []ssh.ClientAuth{
			ssh.ClientAuthKeyring(clientKey),
		},
	}
	addr := fmt.Sprintf("%s:%s", s.Config.Host, s.Config.Port)
	var err error
	s.Client, err = ssh.Dial("tcp", addr, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	s.Connected = true
}

type keychain struct {
	key *rsa.PrivateKey
}

func (k *keychain) Key(i int) (ssh.PublicKey, error) {
	if i != 0 {
		return nil, nil
	}
	return ssh.NewPublicKey(&k.key.PublicKey)
}

func (k *keychain) Sign(i int, rand io.Reader, data []byte) (sig []byte, err error) {
	hashFunc := crypto.SHA1
	h := hashFunc.New()
	h.Write(data)
	digest := h.Sum(nil)
	return rsa.SignPKCS1v15(rand, k.key, hashFunc, digest)
}

func loadPrivateKey(privateKeyPath string) []byte {
	privateKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("Error while reading private key: %s", err)
	}
	return privateKey
}
