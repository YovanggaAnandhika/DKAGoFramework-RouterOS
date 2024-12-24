package ip

import "github.com/go-routeros/routeros/v3"

type IP struct {
	Client *routeros.Client
	Prefix string
}

func (ip IP) print() (reply *routeros.Reply, err error) {
	cmd := ip.Prefix + "/print"
	res, err := ip.Client.Run(cmd)
	return res, err
}
