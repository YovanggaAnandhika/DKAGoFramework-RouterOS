package main

import (
	RouterOS "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS"
	"log"
)

func main() {

	client := RouterOS.Client{
		Host: "server.dkaapis.com",
		Port: 8729,
		Auth: &RouterOS.RouterOSConfigAuth{
			Username: "developer",
			Password: "Cyberhack2010",
		},
		Secure: &RouterOS.RouterOSConfigSecure{
			CACertFile:     "config/ssl/CA.crt",
			ClientCertFile: "config/ssl/RouterOSClient.crt",
			ClientKeyFile:  "config/ssl/RouterOSClient.key",
		},
	}

	connection, err := client.Connect()
	res, err := connection.Interface().Print()

	if err != nil {
		log.Fatalf("error fatal %s", err)
		return
	}
	log.Print(res)
}
