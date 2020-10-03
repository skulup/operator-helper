package certs

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"math/big"
	"net"
	"strings"
	"time"
)

func Generate(KeyType string, hosts string, organization string, country string, years int) ([]byte, []byte, error) {
	// Create the CA certificate template
	caCertTemp := &x509.Certificate{
		IsCA:         true, // set to true to indicate that this is our CA
		SerialNumber: big.NewInt(2020),
		Subject: pkix.Name{
			Organization: []string{organization},
			Country:      toStringArray(country),
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(years, 0, 0),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	// Generate CA public and private key
	caPrivateKey, err := generatePrivateKey(KeyType)
	if err != nil {
		return nil, nil, errors.Wrap(err, "generatePrivateKey")
	}
	// Create the CA certificate and sign it with the keys
	caCert, err := x509.CreateCertificate(rand.Reader,
		caCertTemp, caCertTemp, getPublicKey(caPrivateKey), caPrivateKey)
	if err != nil {
		return nil, nil, errors.Wrap(err, "x509.CreateCertificate")
	}
	// Encode the CA certificate to PEM
	caCertPem := &bytes.Buffer{}
	if err := pem.Encode(caCertPem, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCert,
	}); err != nil {
		return nil, nil, errors.Wrap(err, "CA cert pem.Encode")
	}

	// Generate the server cert template
	serverCertTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization: []string{organization},
			Country:      toStringArray(country),
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(years, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	}
	for _, h := range strings.Split(hosts, ",") {
		if ip := net.ParseIP(h); ip != nil {
			serverCertTemplate.IPAddresses = append(serverCertTemplate.IPAddresses, ip)
		} else {
			serverCertTemplate.DNSNames = append(serverCertTemplate.DNSNames, h)
		}
	}
	// Generate server cert public and private key
	serverCertPrivateKey, err := generatePrivateKey(KeyType)
	if err != nil {
		return nil, nil, errors.Wrap(err, "generatePrivateKey")
	}
	// Create the server cert and sign it with our CA
	serverCert, err := x509.CreateCertificate(rand.Reader, serverCertTemplate, caCertTemp, getPublicKey(serverCertPrivateKey), caPrivateKey)
	if err != nil {
		return nil, nil, err
	}
	// PEM encode the  server cert
	serverCertPEM, err := pemEncodeCertificate(serverCert)
	if err != nil {
		return nil, nil, err
	}
	// PEM encode the  server cert
	serverCertPrivateKeyPEM, err := pemEncodePrivateKey(serverCertPrivateKey, KeyType)
	if err != nil {
		return nil, nil, err
	}
	return serverCertPEM, serverCertPrivateKeyPEM, nil
}

func generatePrivateKey(keyType string) (pv interface{}, err error) {
	switch strings.ToUpper(keyType) {
	case "ECDSA":
		pv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "ED25519":
		_, pv, err = ed25519.GenerateKey(rand.Reader)
	default:
		pv, err = rsa.GenerateKey(rand.Reader, 4096)
	}
	return
}

func getPublicKey(pv interface{}) interface{} {
	switch key := pv.(type) {
	case *rsa.PrivateKey:
		return &key.PublicKey
	case *ecdsa.PrivateKey:
		return &key.PublicKey
	case ed25519.PrivateKey:
		return key.Public().(ed25519.PublicKey)
	default:
		return nil
	}
}

func pemEncodeCertificate(blockData []byte) ([]byte, error) {
	pemBuffer := &bytes.Buffer{}
	if err := pem.Encode(pemBuffer, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: blockData,
	}); err != nil {
		return nil, err
	}
	return pemBuffer.Bytes(), nil
}

func pemEncodePrivateKey(key interface{}, keyType string) ([]byte, error) {
	var blockData []byte
	if _, IsRSA := key.(*rsa.PrivateKey); IsRSA {
		blockData = x509.MarshalPKCS1PrivateKey(key.(*rsa.PrivateKey))
	} else {
		d, err := x509.MarshalPKCS8PrivateKey(key)
		if err != nil {
			return nil, err
		}
		blockData = d
	}
	pemBuffer := &bytes.Buffer{}
	if err := pem.Encode(pemBuffer, &pem.Block{
		Type:  pemBlockType(keyType),
		Bytes: blockData,
	}); err != nil {
		return nil, err
	}
	return pemBuffer.Bytes(), nil
}

func pemBlockType(keyType string) string {
	if keyType == "" {
		keyType = "RSA"
	}
	return fmt.Sprintf("%s PRIVATE KEY", keyType)
}

func toStringArray(s string) []string {
	if s == "" {
		return nil
	}
	return []string{s}
}
