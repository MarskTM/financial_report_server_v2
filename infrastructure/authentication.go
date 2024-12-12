package infrastructure

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"github.com/go-chi/jwtauth"
	"github.com/golang/glog"
)

// Algorithm algorithm define
const Algorithm = "RS256"

func loadAuthToken() error {
	// Load private key
	privateReader, err := ioutil.ReadFile(privatePath)
	if err != nil {
		glog.V(1).Info("NO RSA private pem file")
		return err
	}
	privatePem, _ := pem.Decode(privateReader)

	if privatePem.Type != "RSA PRIVATE KEY" {
		glog.V(1).Info("RSA private key is of the wrong type")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privatePem.Bytes)
	if err != nil {
		glog.Error(err)
	}
	// Read public key
	publicReader, err := ioutil.ReadFile(publicPath)
	if err != nil {
		glog.V(1).Info("No RSA public pem file")
		return err
	}
	publicPem, _ := pem.Decode(publicReader)
	publicKey, _ := x509.ParsePKIXPublicKey(publicPem.Bytes)

	encodeAuth = jwtauth.New(Algorithm, privateKey, publicKey)
	decodeAuth = jwtauth.New(Algorithm, nil, publicKey)

	return nil
}

func RsaEncrypt(decrypt string) (string, error) {
	// Load public key
	publicKey, err := ioutil.ReadFile(publicPath)
	if err != nil {
		glog.V(1).Info("No RSA public pem file")
		return "", err
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		glog.Error(err)
		return "", err
	}
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), []byte(decrypt))
	if err != nil {
		glog.Error(err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func RsaDecrypt(encrypt string) ([]byte, error) {
	// Load private key
	privateKey, err := ioutil.ReadFile(privatePath)
	if err != nil {
		glog.V(1).Info("NO RSA private pem file")
		return nil, err
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	cipherText, _ := base64.StdEncoding.DecodeString(encrypt)
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
}
