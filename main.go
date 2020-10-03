package main

import (
	"fmt"
	"github.com/skulup/operator-helper/certs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := filepath.Join(os.TempDir(), "k8s-webhook-server", "serving-certs")
	cert, key, err := certs.Generate("RSA",
		"pulsar-operator-service.default.svc",
		"Corporation Ltd", "SL", 10)
	if err != nil {
		log.Fatal(err)
	}
	if err = os.MkdirAll(dir, 0777); err != nil {
		log.Fatal(err)
	}
	write := func(filename string, data []byte) {
		path := fmt.Sprintf("%s/%s", dir, filename)
		if err := WriteToFile(path, data); err != nil {
			log.Fatal(err)
		}
	}
	write("tls.crt", cert)
	write("tls.key", key)

}

// WriteFile writes data in the file at the given path
func WriteToFile(filepath string, data []byte) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
