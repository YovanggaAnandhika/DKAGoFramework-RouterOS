package main

import (
	RouterOS "github.com/YovanggaAnandhika/DKAGoFramework-RouterOS"
	"log"
	_ "log"
)

func main() {

	client, _ := RouterOS.Client(&RouterOS.RouterOSConfig{
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
	})

	res, err := client.Interfaces().Print()
	log.Print(res)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Print(res)
}
