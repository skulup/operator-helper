package reconcilers

import (
	"fmt"
	"github.com/skulup/operator-helper/reconciler"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Configure let the added reconcilers to configure themselves
func Configure(manager ctrl.Manager, reconcilers ...reconciler.Reconciler) error {
	ctx := reconciler.NewContext(manager)
	for _, r := range reconcilers {
		fmt.Printf("configuring the reconciler: %T\n", r)
		if err := r.Configure(ctx); err != nil {
			return err
		}
	}
	return nil
}
