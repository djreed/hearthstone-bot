package ssl

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	pool      *x509.CertPool
	client    *http.Client
	tlsConfig *tls.Config
	wsDialer  *websocket.Dialer
)

func init() {
	pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(pemCerts)

	tlsConfig = &tls.Config{RootCAs: pool}

	client = &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}}

	wsDialer = websocket.DefaultDialer
	wsDialer.TLSClientConfig = tlsConfig
}

func HTTPSClient() *http.Client {
	return client
}

func Dialer() *websocket.Dialer {
	return wsDialer
}
