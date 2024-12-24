package action

import (
	"github.com/go-routeros/routeros/v3"
)

type Router struct {
	Client *routeros.Client
	Prefix string
}

func (router *Router) Print() (*routeros.Reply, error) {
	defer func(Client *routeros.Client) {
		err := Client.Close()
		if err != nil {
			return
		}
	}(router.Client)
	//#####################################
	cmd := router.Prefix + "/print"
	//#####################################
	res, err := router.Client.Run(cmd)
	return res, err
}
