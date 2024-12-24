package connection

import "github.com/go-routeros/routeros/v3"

type Connection struct {
	client *routeros.Client
}

func (options Connection) GetInterfaces() {

}
