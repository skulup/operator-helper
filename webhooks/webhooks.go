package webhooks

import (
	"fmt"
	"github.com/skulup/operator-helper/certs"
	"github.com/skulup/operator-helper/configs"
	"k8s.io/apimachinery/pkg/runtime"
	"log"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Configure configures the webhook for the added CR types
func Configure(manager ctrl.Manager, apiTypes ...runtime.Object) error {
	if configs.WebHooksEnabled() {
		if err := generateWebhookCertIfMissing(); err != nil {
			log.Printf("Err creating the default webhook certificates: %s", err)
			return err
		}
		for _, apiType := range apiTypes {
			fmt.Printf("configuring the webhook: %T\n", apiType)
			if err := ctrl.NewWebhookManagedBy(manager).For(apiType).Complete(); err != nil {
				return err
			}
		}
	} else {
		log.Printf("Cannot configure webhooks as it's disabled")
	}
	return nil
}

func generateWebhookCertIfMissing() (err error) {
	dir := configs.GetWebHookCertDir()
	keyPath := fmt.Sprintf("%s/tls.key", dir)
	certPath := fmt.Sprintf("%s/tls.crt", dir)
	if fileDoesNotExists(certPath) || fileDoesNotExists(keyPath) {
		log.Printf("Generating the default webhook certificate to directory: %s", dir)
		hosts := configs.GetWebHookServiceHosts()
		organization := "Apache Pulsar Org"
		cert, key, err0 := certs.Generate("RSA", hosts, organization, "", 10)
		if err0 != nil {
			err = err0
			return
		}
		if err = os.MkdirAll(dir, 0777); err == nil {
			if err = writeToFile(certPath, cert); err == nil {
				err = writeToFile(keyPath, key)
			}
		}
	}
	return
}

func fileDoesNotExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

// WriteFile writes data in the file at the given path
func writeToFile(filepath string, data []byte) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
