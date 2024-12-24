package action

import (
	"github.com/go-routeros/routeros/v3"
)

type Action struct {
	Client *routeros.Client
	Prefix string
}

func (Act *Action) Print() (*routeros.Reply, error) {
	defer func(Client *routeros.Client) {
		err := Client.Close()
		if err != nil {
			return
		}
	}(Act.Client)
	//#####################################
	cmd := Act.Prefix + "/print"
	//#####################################
	res, err := Act.Client.Run(cmd)
	return res, err
}
