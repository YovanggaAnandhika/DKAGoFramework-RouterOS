package interfaces

import "github.com/go-routeros/routeros/v3"

type Interfaces struct {
	Client *routeros.Client
	Prefix string
}

func (interfaces *Interfaces) Print() (reply *routeros.Reply, err error) {
	cmd := interfaces.Prefix + "/print"
	res, err := interfaces.Client.Run(cmd)
	return res, err
}
