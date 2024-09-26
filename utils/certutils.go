package utils

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// paths 证书路径, 证书格式为PEM
func LoadCert(paths ...string) *x509.CertPool {

	CAs, err := x509.SystemCertPool()
	if err != nil {
		log.Errorf("load system cert error. %v", err)
	}

	if CAs == nil {
		CAs = x509.NewCertPool()
	}

	// append cert to pool
	for _, path := range paths {
		pem, err := os.ReadFile(path)
		if err != nil {
			log.Errorf("read pem file error. %v", err)
		}

		CAs.AppendCertsFromPEM(pem)
	}

	return CAs
}

func HttpClientWithCACerts(certPath string) *http.Client {
	CAs := LoadCert(certPath)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            CAs,
			ServerName:         "", // 对应证书中的 DNS name; 当对端证书的DNS name和 访问的URL中的hostname不一致时,可以通过设置此 serverName来通过验证
		},
	}

	client := &http.Client{
		Timeout:   50 * time.Second,
		Transport: tr,
	}

	return client
}
