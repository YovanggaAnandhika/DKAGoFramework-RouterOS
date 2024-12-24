package firewall

import (
	filter "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/ip/firewall/Filter"
	mangle "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/ip/firewall/Mangle"
	nat "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/ip/firewall/Nat"
	"github.com/go-routeros/routeros/v3"
)

type Firewall struct {
	Client *routeros.Client
	Prefix string
}

func (Firewall *Firewall) Filter() *filter.Filter {
	//#############################################
	prefix := Firewall.Prefix + "/filter"
	//#############################################
	return &filter.Filter{
		Client: Firewall.Client,
		Prefix: prefix,
	}
}

func (Firewall *Firewall) Nat() *nat.Nat {
	//#############################################
	prefix := Firewall.Prefix + "/nat"
	//#############################################
	return &nat.Nat{
		Client: Firewall.Client,
		Prefix: prefix,
	}
}

func (Firewall *Firewall) Mangle() *mangle.Mangle {
	//#############################################
	prefix := Firewall.Prefix + "/nat"
	//#############################################
	return &mangle.Mangle{
		Client: Firewall.Client,
		Prefix: prefix,
	}
}
