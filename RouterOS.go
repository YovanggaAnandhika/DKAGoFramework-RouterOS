package RouterOS

import (
	"crypto/tls"
	"github.com/go-routeros/routeros/v3"
	"log"
	"strconv"
)

type RouterOSConfigAuth struct {
	Username *string
	Password *string
}

type RouterOSConfig struct {
	Host   *string
	Port   *int
	Auth   *RouterOSConfigAuth
	Secure *tls.Config
}

// Default values for RouterOSConfig
var (
	defaultHost     = "192.168.88.1"
	defaultPort     = 8728
	defaultUsername = "admin"
	defaultPassword = ""
	defaultSecure   = &tls.Config{InsecureSkipVerify: true}
)

// Connect establishes a connection to a RouterOS device and returns the client instance.
func (options *RouterOSConfig) Connect() (*routeros.Client, error) {
	// Set default values if not provided
	if options.Host == nil {
		host := defaultHost
		options.Host = &host
	}
	if options.Port == nil {
		port := defaultPort
		options.Port = &port
	}
	if options.Auth == nil {
		options.Auth = &RouterOSConfigAuth{}
	}
	if options.Auth.Username == nil {
		username := defaultUsername
		options.Auth.Username = &username
	}
	if options.Auth.Password == nil {
		password := defaultPassword
		options.Auth.Password = &password
	}
	if options.Secure == nil {
		options.Secure = defaultSecure
	}

	// Construct address
	address := *options.Host + ":" + strconv.Itoa(*options.Port)

	var client *routeros.Client
	var err error

	// Connect using TLS if secure config is provided, otherwise use plain connection
	if options.Secure != nil {
		client, err = routeros.DialTLS(address, *options.Auth.Username, *options.Auth.Password, options.Secure)
	} else {
		client, err = routeros.Dial(address, *options.Auth.Username, *options.Auth.Password)
	}

	if err != nil {
		log.Printf("Failed to connect to RouterOS: %v", err)
		return nil, err
	}

	return client, nil
}

// GetInstance creates and returns a RouterOS client instance using the provided configuration.
func Client(options *RouterOSConfig) (*routeros.Client, error) {
	client, err := options.Connect()
	if err != nil {
		return nil, err
	}
	return client, nil
}
