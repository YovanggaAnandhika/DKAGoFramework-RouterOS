package RouterOS

import (
	"crypto/tls"
	"github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/connection"
	"github.com/go-routeros/routeros/v3"
	"log"
	"strconv"
)

type RouterOSConfigAuth struct {
	Username string
	Password string
}

type RouterOSConfig struct {
	Host   string
	Port   int
	Auth   *RouterOSConfigAuth
	Secure *tls.Config
}

// Default values for RouterOSConfig
var (
	defaultHost     = "192.168.1.1"
	defaultPort     = 8729
	defaultUsername = "admin"
	defaultPassword = ""
	defaultSecure   = &tls.Config{}
)

// Connect establishes a connection to a RouterOS device and returns the client instance.
func Client(options *RouterOSConfig) (*connection.Connection, error) {
	// Set default values if not provided
	if options.Host == "" {
		options.Host = defaultHost
	}
	if options.Port == 0 {
		options.Port = defaultPort
	}
	if options.Auth == nil {
		options.Auth = &RouterOSConfigAuth{}
	}
	if options.Auth.Username == "" {
		options.Auth.Username = defaultUsername
	}
	if options.Auth.Password == "" {
		options.Auth.Password = defaultPassword
	}
	if options.Secure == nil {
		options.Secure = defaultSecure
	}

	// Construct address
	address := options.Host + ":" + strconv.Itoa(options.Port)

	var client *routeros.Client
	var err error

	// Connect using TLS if secure config is provided, otherwise use plain connection
	if options.Secure != nil {
		client, err = routeros.DialTLS(address, options.Auth.Username, options.Auth.Password, options.Secure)
	} else {
		client, err = routeros.Dial(address, options.Auth.Username, options.Auth.Password)
	}

	if err != nil {
		log.Printf("Failed to connect to RouterOS: %v", err)
		return nil, err
	}

	return &connection.Connection{client: client}, nil
}
