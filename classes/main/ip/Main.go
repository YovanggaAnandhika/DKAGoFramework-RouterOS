package ip

import (
	"github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/ip/address"
	"github.com/go-routeros/routeros/v3"
)

type IP struct {
	Client *routeros.Client
	Prefix string
}

func (IP *IP) Address() *action.Action {
	//#############################################
	prefix := IP.Prefix + "/address"
	//#############################################
	return &action.Action{
		Client: IP.Client,
		Prefix: prefix,
	}
}
