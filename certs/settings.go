package certs

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

func HTTPSClient() *http.Client {
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(pemCerts)
	return &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool}}}
}
