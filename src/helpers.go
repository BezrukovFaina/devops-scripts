package devops_scripts

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GenerateKeyPair generates a new RSA key pair and saves it to a file
func GenerateKeyPair(privateKeyPath, publicKeyPath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err!= nil {
		return err
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes},
	)

	publicKeyBytes := &privateKey.PublicKey
	publicKeyPEM := x509.MarshalPKIXPublicKey(publicKeyBytes)
	publicKeyPEM = pem.EncodeToMemory(
		&pem.Block{Type: "RSA PUBLIC KEY", Bytes: publicKeyPEM},
	)

	if err := ioutil.WriteFile(privateKeyPath, privateKeyPEM, 0600); err!= nil {
		return err
	}

	if err := ioutil.WriteFile(publicKeyPath, publicKeyPEM, 0644); err!= nil {
		return err
	}

	return nil
}

// LoadPublicKey loads a public key from a file and returns it
func LoadPublicKey(keyPath string) (*rsa.PublicKey, error) {
	data, err := ioutil.ReadFile(keyPath)
	if err!= nil {
		return nil, err
	}

	var publicKey interface{}
	if err := pem.Decode(data); err!= nil {
		return nil, err
	}

	switch publicKey := publicKey.(type) {
	case *pem.Block:
		if publicKey.Type!= "RSA PUBLIC KEY" {
			return nil, fmt.Errorf("unknown key type: %s", publicKey.Type)
		}

		publicKeyBytes, ok := publicKey.Bytes.([]byte)
		if!ok {
			return nil, fmt.Errorf("invalid key data")
		}

		publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
		if err!= nil {
			return nil, err
		}

		return publicKey.(*rsa.PublicKey), nil
	default:
		return nil, fmt.Errorf("unknown key type: %T", publicKey)
	}
}

// LoadPrivateKey loads a private key from a file and returns it
func LoadPrivateKey(keyPath string) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(keyPath)
	if err!= nil {
		return nil, err
	}

	var privateKey interface{}
	if err := pem.Decode(data); err!= nil {
		return nil, err
	}

	switch privateKey := privateKey.(type) {
	case *pem.Block:
		if privateKey.Type!= "RSA PRIVATE KEY" {
			return nil, fmt.Errorf("unknown key type: %s", privateKey.Type)
		}

		privateKeyBytes, ok := privateKey.Bytes.([]byte)
		if!ok {
			return nil, fmt.Errorf("invalid key data")
		}

		privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
		if err!= nil {
			return nil, err
		}

		return privateKey, nil
	default:
		return nil, fmt.Errorf("unknown key type: %T", privateKey)
	}
}

// GetEnvironmentVariable returns the value of an environment variable
func GetEnvironmentVariable(name string) string {
	return os.Getenv(name)
}

// GetHomeDirectory returns the user's home directory
func GetHomeDirectory() string {
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE")
	}
	return home
}

// GetConfigPath returns the path to the project's config file
func GetConfigPath() string {
	return filepath.Join(GetHomeDirectory(), ".config", "devops-scripts")
}

// GetPublicKeyPath returns the path to the project's public key file
func GetPublicKeyPath() string {
	return filepath.Join(GetConfigPath(), "public_key.pem")
}

// GetPrivateKeyPath returns the path to the project's private key file
func GetPrivateKeyPath() string {
	return filepath.Join(GetConfigPath(), "private_key.pem")
}