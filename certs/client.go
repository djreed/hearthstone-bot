package certs

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

var (
	pool   *x509.CertPool
	client *http.Client
)

func init() {
	pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(pemCerts)
	client = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool}}}
}

func HTTPSClient() *http.Client {
	return client
}
