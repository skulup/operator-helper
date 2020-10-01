package webhooks

import (
	"fmt"
	"github.com/wiretld/operator-pkg/configs"
	"k8s.io/apimachinery/pkg/runtime"
	"log"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Configure configures the webhook for the added CR types
func Configure(manager ctrl.Manager, apiTypes ...runtime.Object) error {
	if configs.WebHooksEnabled() {
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
