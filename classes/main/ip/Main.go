package ip

import (
	address "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/ip/address"
	"github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/ip/firewall"
	"github.com/go-routeros/routeros/v3"
)

type IP struct {
	Client *routeros.Client
	Prefix string
}

func (IP *IP) Address() *address.Action {
	//#############################################
	prefix := IP.Prefix + "/address"
	//#############################################
	return &address.Action{
		Client: IP.Client,
		Prefix: prefix,
	}
}

func (IP *IP) Firewall() *firewall.Firewall {
	//#############################################
	prefix := IP.Prefix + "/firewall"
	//#############################################
	return &firewall.Firewall{
		Client: IP.Client,
		Prefix: prefix,
	}
}
