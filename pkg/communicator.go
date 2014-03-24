package pkg

type Communicator interface {
	Run(command string) string
}
