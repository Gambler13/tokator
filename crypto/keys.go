package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	v1 "k8s.io/api/core/v1"
)

const BlockTypePrivate = "RSA PRIVATE KEY"
const BlockTypePublic = "RSA PUBLIC KEY"
const SecretKeyPrivateKey = "private_key"

func decodeRSA(key *rsa.PrivateKey) (private string, public string) {
	public = string(pem.EncodeToMemory(
		&pem.Block{
			Type:  BlockTypePublic,
			Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey),
		},
	))

	private = string(pem.EncodeToMemory(
		&pem.Block{
			Type:  BlockTypePrivate,
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	))

	return
}

func DecodeRSAPublic(key rsa.PublicKey) (public string) {
	public = string(pem.EncodeToMemory(
		&pem.Block{
			Type:  BlockTypePublic,
			Bytes: x509.MarshalPKCS1PublicKey(&key),
		},
	))

	return
}

func EncodePublicRSA(public string) (*rsa.PublicKey, error) {

	block, _ := pem.Decode([]byte(public))
	if block == nil || block.Type != BlockTypePublic {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	return x509.ParsePKCS1PublicKey(block.Bytes)
}

func encodeRSA(private string) (*rsa.PrivateKey, error) {

	block, _ := pem.Decode([]byte(private))
	if block == nil || block.Type != BlockTypePrivate {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func FromSecret(secret *v1.Secret) (*rsa.PrivateKey, error) {
	privatePem := secret.Data[SecretKeyPrivateKey]
	return encodeRSA(string(privatePem))
}

func ToSecret(key *rsa.PrivateKey, secret *v1.Secret) {
	priv, _ := decodeRSA(key)
	DecodedToSecret(priv, secret)
}

func DecodedToSecret(private string, secret *v1.Secret) {
	secret.StringData = map[string]string{SecretKeyPrivateKey: private}
}

func CreateKeys() (private, public string, err error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	private, public = decodeRSA(key)
	return
}
