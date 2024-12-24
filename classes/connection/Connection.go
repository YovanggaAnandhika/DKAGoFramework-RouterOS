package connection

import "github.com/go-routeros/routeros/v3"

type Connection struct {
	Client *routeros.Client
}

func (options Connection) GetInterfaces() {

}
