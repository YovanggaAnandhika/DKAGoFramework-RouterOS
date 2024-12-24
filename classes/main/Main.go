package classes

import (
	iface "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/interface"
	ip "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/ip"
	system "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/system"
	user_manager "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/user-manager"
	"github.com/go-routeros/routeros/v3"
)

type Classes struct {
	Client *routeros.Client
}

func (options *Classes) UserManager() *user_manager.UserManager {
	//#############################################
	prefix := "/user-manager"
	//##############################################
	//#############################################
	return &user_manager.UserManager{
		Client: options.Client,
		Prefix: prefix,
	}
}

func (options *Classes) Interface() *iface.Interfaces {
	//#############################################
	prefix := "/interface"
	//##############################################
	//#############################################
	return &iface.Interfaces{
		Client: options.Client,
		Prefix: prefix,
	}
}

func (options *Classes) Ip() *ip.IP {
	//#############################################
	prefix := "/ip"
	//#############################################
	return &ip.IP{
		Client: options.Client,
		Prefix: prefix,
	}
}

func (options *Classes) System() *system.SYSTEM {
	//#############################################
	prefix := "/system"
	//#############################################
	return &system.SYSTEM{
		Client: options.Client,
		Prefix: prefix,
	}
}
