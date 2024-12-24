package action

import (
	"github.com/go-routeros/routeros/v3"
)

type Nat struct {
	Client *routeros.Client
	Prefix string
}

func (nat *Nat) Print() (*routeros.Reply, error) {
	defer func(Client *routeros.Client) {
		err := Client.Close()
		if err != nil {
			return
		}
	}(nat.Client)
	//#####################################
	cmd := nat.Prefix + "/print"
	//#####################################
	res, err := nat.Client.Run(cmd)
	return res, err
}
