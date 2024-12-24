package system

import (
	resource "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/system/resource"
	"github.com/go-routeros/routeros/v3"
)

type SYSTEM struct {
	Client *routeros.Client
	Prefix string
}

func (System *SYSTEM) Resource() *resource.Resource {
	//#############################################
	prefix := System.Prefix + "/resource"
	//#############################################
	return &resource.Resource{
		Client: System.Client,
		Prefix: prefix,
	}
}
