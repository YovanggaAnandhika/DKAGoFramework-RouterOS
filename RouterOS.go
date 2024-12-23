package RouterOS

import (
	"crypto/tls"
	"github.com/go-routeros/routeros/v3"
	"log"
	"strconv"
)

type RouterOSConfigAuth struct {
	username *string
	password *string
}

type RouterOSConfig struct {
	host   *string
	port   *int
	auth   *RouterOSConfigAuth
	secure *tls.Config
}

// Default values for RouterOSConfig
var (
	DefaultHost     = "192.168.88.1"
	DefaultPort     = 8728
	DefaultUsername = "admin"
	DefaultPassword = ""
	DefaultSecure   = &tls.Config{InsecureSkipVerify: true}
)

// getInstance returns a connected RouterOS client instance.
func (options *RouterOSConfig) getInstance() *routeros.Client {
	// Set default values if not provided
	if options.host == nil {
		host := DefaultHost
		options.host = &host
	}
	if options.port == nil {
		port := DefaultPort
		options.port = &port
	}
	if options.auth == nil {
		options.auth = &RouterOSConfigAuth{}
	}
	if options.auth.username == nil {
		username := DefaultUsername
		options.auth.username = &username
	}
	if options.auth.password == nil {
		password := DefaultPassword
		options.auth.password = &password
	}
	if options.secure == nil {
		options.secure = DefaultSecure
	}

	// Construct address
	address := *options.host + ":" + strconv.Itoa(*options.port)
	var client *routeros.Client
	var err error

	if options.secure != nil {
		client, err = routeros.DialTLS(address, *options.auth.username, *options.auth.password, options.secure)
	} else {
		client, err = routeros.Dial(address, *options.auth.username, *options.auth.password)
	}

	if err != nil {
		log.Fatalf("Failed to connect to RouterOS: %v", err)
	}

	// Ensure the connection is closed when done
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	return client
}
