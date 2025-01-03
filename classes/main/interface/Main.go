package _interface

import (
	"github.com/go-routeros/routeros/v3"
)

type Interfaces struct {
	Client *routeros.Client
	Prefix string
}

func (interfaces *Interfaces) Print() (*routeros.Reply, error) {
	defer func(Client *routeros.Client) {
		err := Client.Close()
		if err != nil {
			return
		}
	}(interfaces.Client)
	//#####################################
	cmd := interfaces.Prefix + "/print"
	//#####################################
	res, err := interfaces.Client.Run(cmd)
	if err != nil {
		return nil, err
	}
	return res, nil
}
