package cryptoutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func ExportPubKeyAsPEMStr(pubkey *rsa.PublicKey) string {
	pubKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pubkey),
		},
	))
	return pubKeyPem
}

func ExportPrivKeyAsPEMStr(privkey *rsa.PrivateKey) string {
	privKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privkey),
		},
	))
	return privKeyPem
}

func ImportPEMStrAsPrivateKey(privateKey string) (int, *rsa.PrivateKey) {
	privateKeyPEM := []byte(privateKey)
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		return -1, nil
	}
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return -1, nil
	}
	return 1, privKey
}

func ImportPEMStrAsPublicKey(publicKey string) (int, *rsa.PublicKey) {
	publickKeyPEM := []byte(publicKey)
	block, _ := pem.Decode(publickKeyPEM)
	if block == nil {
		return -1, nil
	}
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return -1, nil
	}
	return 1, pubKey
}

func RSA_GenKey(privateKey *string, publicKey *string, errDesc *string) int {
	privKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		*errDesc = fmt.Sprintf("rsa.GenerateKey() failed with err:%s", err)
		return -1
	}
	*privateKey = ExportPrivKeyAsPEMStr(privKey)
	*publicKey = ExportPubKeyAsPEMStr(&privKey.PublicKey)
	return 1
}

func RSA_OAEP_Encrypt(publicKey string, inputData string, eData *string, errDesc *string, paddingLabel string) int {
	rVal, pubKey := ImportPEMStrAsPublicKey(publicKey)
	if rVal < 0 {
		*errDesc = "importPEMStrAsPublicKey failed"
		return -1
	}
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, []byte(inputData), []byte(paddingLabel))
	if err != nil {
		*errDesc = fmt.Sprintf("rsa.DecryptOAEP() failed with err:%s", err)
		return -1
	}
	*eData = base64.StdEncoding.EncodeToString(ciphertext)
	return 1
}

func RSA_OAEP_Decrypt(privateKey string, eData string, clearData *string, errDesc *string, paddingLabel string) int {
	ct, _ := base64.StdEncoding.DecodeString(eData)
	rVal, privKey := ImportPEMStrAsPrivateKey(privateKey)
	if rVal < 0 {
		*errDesc = "importPEMStrAsPrivateKey failed"
		return -1
	}
	decryptedPlaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, []byte(ct), []byte(paddingLabel))
	if err != nil {
		*errDesc = fmt.Sprintf("privKey.Decrypt failed with err:%s", err)
		return -1
	}
	*clearData = string(decryptedPlaintext)
	return 1
}
