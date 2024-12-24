package interfaces

import (
	"github.com/go-routeros/routeros/v3"
)

type Interfaces struct {
	Client *routeros.Client
	Prefix string
}

func (interfaces *Interfaces) Print() (*routeros.Reply, error) {
	cmd := interfaces.Prefix + "/print"
	res, err := interfaces.Client.Run(cmd)
	if err != nil {
		return nil, err
	}
	return res, nil
}
