package main

/*

Example:

  go run  generate_cert.go -email-address=a@a.com -host localhost,napoleon.lan,104.236.87.120,104.197.33.112 -ca  -ecdsa-curve P521




 */


import (
	"crypto/tls"
	"crypto/x509"
	"github.com/levigross/grequests"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Load our TLS key pair to use for authentication
	cert, err := tls.LoadX509KeyPair("/Users/mchirico/testCert/cert.pem",
		"/Users/mchirico/testCert/key.pem")
	if err != nil {
		log.Fatalln("Unable to load cert", err)
	}

	// Load our CA certificate
	clientCACert, err := ioutil.ReadFile("/Users/mchirico/testCert/cert.pem")
	if err != nil {
		log.Fatal("Unable to open cert", err)
	}

	clientCertPool := x509.NewCertPool()
	clientCertPool.AppendCertsFromPEM(clientCACert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      clientCertPool,
	}

	tlsConfig.BuildNameToCertificate()


	ro := &grequests.RequestOptions{
		HTTPClient: &http.Client{
			Transport: &http.Transport{TLSClientConfig: tlsConfig},
		},
	}
	resp, err := grequests.Get("https://localhost:8080", ro)
	if err != nil {
		log.Println("Unable to speak to our server", err)
	}

	// Lets print the message
	log.Println(resp.String())
}
