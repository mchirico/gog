package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*


I had to modify generate_cert.go, to include email-address.


*/

/*

Example:

  go run  generate_cert.go -email-address=a@a.com -host localhost,napoleon.lan,104.236.87.120,104.197.33.112 -ca  -ecdsa-curve P521




 */

func main() {

	caCert, err := ioutil.ReadFile("/Users/mchirico/testCert/cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		// Reject any TLS certificate that cannot be validated
		ClientAuth: tls.RequireAndVerifyClientCert,
		// Ensure that we only use our "CA" to validate certificates
		ClientCAs: caCertPool,

		CipherSuites: []uint16{tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256},
		// Force it server side
		PreferServerCipherSuites: true,
		// TLS 1.2 because we can
		MinVersion: tls.VersionTLS12,
	}

	tlsConfig.BuildNameToCertificate()

	http.HandleFunc("/", HelloUser)

	httpServer := &http.Server{
		Addr:      ":8080",
		TLSConfig: tlsConfig,
	}

	log.Println(httpServer.ListenAndServeTLS("/Users/mchirico/testCert/cert.pem",
		"/Users/mchirico/testCert/key.pem"))

}

func HelloUser(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello %v, DNS %v IPAddress: %v\n", req.TLS.PeerCertificates[0].EmailAddresses[0],
		req.TLS.PeerCertificates[0].DNSNames, req.TLS.PeerCertificates[0].IPAddresses)
}
