package acmeaccount

import (
	"os"
	"strings"
	"github.com/go-acme/lego/v3/certcrypto"
	"log"
)
const filePerm os.FileMode = 0600

func createNonExistingFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0700)
	} else if err != nil {
		return err
	}
	return nil
}
func getKeyType(keyType string) certcrypto.KeyType {
	switch strings.ToUpper(keyType) {
	case "RSA2048":
		return certcrypto.RSA2048
	case "RSA4096":
		return certcrypto.RSA4096
	case "RSA8192":
		return certcrypto.RSA8192
	case "EC256":
		return certcrypto.EC256
	case "EC384":
		return certcrypto.EC384
	}

	log.Fatalf("Unsupported KeyType: %s", keyType)
	return ""
}