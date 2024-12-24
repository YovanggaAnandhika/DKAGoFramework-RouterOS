package ip

import "github.com/go-routeros/routeros/v3"

type IP struct {
	Client *routeros.Client
	Prefix string
}

func (IP *IP) Print() (*routeros.Reply, error) {
	cmd := IP.Prefix + "/print"
	res, err := IP.Client.Run(cmd)
	return res, err
}
