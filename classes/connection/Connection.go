package connection

import (
	"github.com/go-routeros/routeros/v3"
	"log"
)

type Connection struct {
	Client *routeros.Client
}

func (options Connection) GetInterfaces() (res *routeros.Reply, err error) {
	defer func(Client *routeros.Client) {
		err := Client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(options.Client)
	//#############################################
	cmd := "/interface/print"
	//#############################################
	res, err = options.Client.Run(cmd)

	return res, err
}
