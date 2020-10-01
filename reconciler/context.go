package reconciler

import (
	"context"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

var (
	_ Context = &contextImpl{}
)

// NewContext creates a new reconciler Context
func NewContext(mgr manager.Manager) Context {
	return &contextImpl{manager: mgr}
}

type contextImpl struct {
	manager manager.Manager
}

func (c *contextImpl) Logger() logr.Logger {
	return c.manager.GetLogger()
}

func (c *contextImpl) Client() client.Client {
	return c.manager.GetClient()
}

func (c *contextImpl) Scheme() *runtime.Scheme {
	return c.manager.GetScheme()
}

func (c *contextImpl) NewControllerBuilder() *builder.Builder {
	return ctrl.NewControllerManagedBy(c.manager)
}

func (c *contextImpl) SetOwnershipReference(owner metav1.Object, controlled metav1.Object) error {
	return controllerutil.SetControllerReference(owner, controlled, c.Scheme())
}

func (c *contextImpl) GetResource(key client.ObjectKey, object runtime.Object, foundCallback func() (err error), notFoundCallback func() (err error)) (err error) {
	err = c.Client().Get(context.TODO(), key, object)
	if err == nil && foundCallback != nil {
		return foundCallback()
	} else if err != nil && errors.IsNotFound(err) {
		return notFoundCallback()
	}
	return
}

func (c *contextImpl) Run(req reconcile.Request, object KubeRuntimeObject, reconcile func() error) (reconcile.Result, error) {
	startTime := time.Now()
	start(req, c.Logger())
	defer stop(req, startTime, c.Logger())
	if err := c.Client().Get(context.TODO(), req.NamespacedName, object); err != nil {
		if errors.IsNotFound(err) {
			// The runtime object is not found. Kubernetes will automatically
			// garbage collect all owned resources - return but do not requeue
			return complete(req, c.Logger())
		}
		// Read error; requeue here
		return errored(err, req, c.Logger())
	}
	if delTime := object.GetDeletionTimestamp(); delTime != nil {
		c.Logger().Info("The request object has been scheduled for delete",
			"Timestamp", time.Until((*delTime).Time).Seconds())
		// The runtime object is marked for deletion - return but do not requeue
		return complete(req, c.Logger())
	}

	if df, ok := object.(Defaulting); ok && df.SetSpecDefaults() {
		c.Logger().Info("Setting the default spec of the request object")
		if err := c.Client().Update(context.TODO(), object); err != nil {
			return errored(err, req, c.Logger())
		}
		return requeue(req, c.Logger())
	}
	if df, ok := object.(Defaulting); ok && df.SetStatusDefaults() {
		c.Logger().Info("Setting the default status of the request object")
		if err := c.Client().Status().Update(context.TODO(), object); err != nil {
			return errored(err, req, c.Logger())
		}
		return requeue(req, c.Logger())
	}
	if err := reconcile(); err != nil {
		return errored(err, req, c.Logger())
	}
	return complete(req, c.Logger())
}

func requeue(req reconcile.Request, logger logr.Logger) (ctrl.Result, error) {
	logger.Info("[Requeue] Reconciliation",
		"Request.Namespace",
		req.NamespacedName, "Request.Name",
		req.Name,
	)
	return reconcile.Result{Requeue: true}, nil
}

func complete(req reconcile.Request, logger logr.Logger) (reconcile.Result, error) {
	logger.Info("[Complete] Reconciliation",
		"Request.Namespace",
		req.NamespacedName, "Request.Name",
		req.Name,
	)
	return reconcile.Result{Requeue: false}, nil
}

func errored(err error, req reconcile.Request, logger logr.Logger) (reconcile.Result, error) {
	logger.Error(err, "[Error] Reconciliation",
		"Request.Namespace",
		req.NamespacedName, "Request.Name",
		req.Name,
	)
	return reconcile.Result{Requeue: true}, err
}

func start(req reconcile.Request, logger logr.Logger) {
	logger.Info("[Start] Reconciliation",
		"Request.Namespace",
		req.Namespace, "Request.Name",
		req.Name,
	)
}

func stop(req reconcile.Request, startTime time.Time, logger logr.Logger) {
	duration := time.Since(startTime).Seconds()
	logger.Info("[Stop] Reconciliation",
		"Request.Namespace",
		req.NamespacedName, "Request.Name",
		req.Name, "durationSec", duration,
	)
}
