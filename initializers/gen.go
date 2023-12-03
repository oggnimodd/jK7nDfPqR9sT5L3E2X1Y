package initializers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func Gen() {
	// Generate a new private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert the private key to PKCS#1 ASN.1 DER encoded form.
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	// Create a new file for the private key.
	privateKeyFile, err := os.Create("private.pem")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write the private key to the file.
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	pem.Encode(privateKeyFile, privateKeyPEM)

	// Generate the public key from the private key.
	publicKey := privateKey.PublicKey

	// Convert the public key to PKIX ASN.1 DER encoded form.
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a new file for the public key.
	publicKeyFile, err := os.Create("public.pem")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write the public key to the file.
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	pem.Encode(publicKeyFile, publicKeyPEM)
}
