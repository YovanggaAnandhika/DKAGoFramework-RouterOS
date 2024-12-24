package RouterOS

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	classes "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS/classes/main"
	"github.com/go-routeros/routeros/v3"
	"os"
	"strconv"
)

type RouterOSConfigAuth struct {
	Username string
	Password string
}

type RouterOSConfigSecure struct {
	CACertFile     string
	ClientCertFile string
	ClientKeyFile  string
}

type Client struct {
	Host   string
	Port   int
	Auth   *RouterOSConfigAuth
	Secure *RouterOSConfigSecure
}

// Default values for RouterOSConfig
var (
	defaultHost     = "192.168.1.1"
	defaultPort     = 8729
	defaultUsername = "admin"
	defaultPassword = ""
	defaultSecure   = &RouterOSConfigSecure{}
)

// Function to configure TLS with options
func createTLSConfig(certFile string, keyFile string, caCertFile string) (*tls.Config, error) {
	// Load the client certificate
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load client certificate: %v", err)
	}

	// Load CA certificate
	caCert, err := os.ReadFile(caCertFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %v", err)
	}

	// Create a Certificate Pool for the CA
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS config with custom options
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false, // Ensure server certificate verification (set to true to skip)
		RootCAs:            caCertPool,
		Certificates:       []tls.Certificate{certificate}, // Set the client certificate
	}

	return tlsConfig, nil
}

// Connect Client Connect establishes a connection to a RouterOS device and returns the client instance.
func (options *Client) Connect() (*classes.Classes, error) {

	var client *routeros.Client
	var err error

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

	TLSConfig, err := createTLSConfig(options.Secure.ClientCertFile, options.Secure.ClientKeyFile, options.Secure.CACertFile)

	// Construct address
	address := options.Host + ":" + strconv.Itoa(options.Port)

	// Connect using TLS if secure config is provided, otherwise use plain connection
	if options.Secure != nil {
		client, err = routeros.DialTLS(address, options.Auth.Username, options.Auth.Password, TLSConfig)
	} else {
		client, err = routeros.Dial(address, options.Auth.Username, options.Auth.Password)
	}

	return &classes.Classes{Client: client}, err
}
