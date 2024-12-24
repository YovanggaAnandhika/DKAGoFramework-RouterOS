package classes

import (
	"github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/interfaces"
	"github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main/ip"
	"github.com/go-routeros/routeros/v3"
	"log"
)

type Classes struct {
	Client *routeros.Client
}

func (options Classes) Interfaces() *interfaces.Interfaces {
	defer func(Client *routeros.Client) {
		err := Client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(options.Client)
	//#############################################
	cmd := "/interface"
	//#############################################
	return &interfaces.Interfaces{
		Client: options.Client,
		Prefix: cmd,
	}
}

func (options Classes) Ip() *ip.IP {
	defer func(Client *routeros.Client) {
		err := Client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(options.Client)
	//#############################################
	cmd := "/ip"
	//#############################################
	return &ip.IP{
		Client: options.Client,
		Prefix: cmd,
	}
}
