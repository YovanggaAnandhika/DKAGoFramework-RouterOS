package user_manager

import (
	router "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/user-manager/Router"
	"github.com/go-routeros/routeros/v3"
)

type UserManager struct {
	Client *routeros.Client
	Prefix string
}

func (usermanager *UserManager) Router() *router.Router {
	//#############################################
	prefix := usermanager.Prefix + "/router"
	//#############################################
	return &router.Router{
		Client: usermanager.Client,
		Prefix: prefix,
	}
}
