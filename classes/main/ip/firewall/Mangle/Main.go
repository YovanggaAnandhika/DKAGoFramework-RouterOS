package action

import (
	"github.com/go-routeros/routeros/v3"
)

type Mangle struct {
	Client *routeros.Client
	Prefix string
}

func (mangle *Mangle) Print() (*routeros.Reply, error) {
	defer func(Client *routeros.Client) {
		err := Client.Close()
		if err != nil {
			return
		}
	}(mangle.Client)
	//#####################################
	cmd := mangle.Prefix + "/print"
	//#####################################
	res, err := mangle.Client.Run(cmd)
	return res, err
}
