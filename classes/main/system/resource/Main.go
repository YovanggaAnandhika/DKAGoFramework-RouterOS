package action

import (
	"github.com/go-routeros/routeros/v3"
)

type Resource struct {
	Client *routeros.Client
	Prefix string
}

func (resource *Resource) Print() (*routeros.Reply, error) {
	defer func(Client *routeros.Client) {
		err := Client.Close()
		if err != nil {
			return
		}
	}(resource.Client)
	//#####################################
	cmd := resource.Prefix + "/print"
	//#####################################
	res, err := resource.Client.Run(cmd)
	return res, err
}
