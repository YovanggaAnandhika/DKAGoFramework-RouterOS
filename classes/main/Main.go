package classes

import (
	"github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/interfaces"
	"github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/ip"
	"github.com/go-routeros/routeros/v3"
)

type Classes struct {
	Client *routeros.Client
}

func (options *Classes) Interface() *interfaces.Interfaces {
	//#############################################
	prefix := "/interface"
	//##############################################
	//#############################################
	return &interfaces.Interfaces{
		Client: options.Client,
		Prefix: prefix,
	}
}

func (options *Classes) Ip() *ip.IP {

	defer func(Client *routeros.Client) {
		err := Client.Close()
		if err != nil {
			return
		}
	}(options.Client)
	//#############################################
	prefix := "/ip"
	//#############################################
	return &ip.IP{
		Client: options.Client,
		Prefix: prefix,
	}
}
