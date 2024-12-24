package action

import (
	"github.com/go-routeros/routeros/v3"
)

type Filter struct {
	Client *routeros.Client
	Prefix string
}

func (filter *Filter) Print() (*routeros.Reply, error) {
	defer func(Client *routeros.Client) {
		err := Client.Close()
		if err != nil {
			return
		}
	}(filter.Client)
	//#####################################
	cmd := filter.Prefix + "/print"
	//#####################################
	res, err := filter.Client.Run(cmd)
	return res, err
}
