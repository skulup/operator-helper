package reconciler

import (
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Reconciler presents the interface to be
// implemented by a controller-runtime controller
type Reconciler interface {
	reconcile.Reconciler

	// Configure configures the reconciler
	Configure(ctx Context) error
}

// Defaulting defines interface for the kubernetes object that provides default spec and status
type Defaulting interface {
	runtime.Object
	metav1.Object

	// Set the default of the object spec and returns true if any set otherwise false
	SetSpecDefaults() bool

	// Set the default of the object status and returns true if any set otherwise false
	SetStatusDefaults() bool
}

// KubeRuntimeObject defines interface of the kubernetes object to reconcile
type KubeRuntimeObject interface {
	runtime.Object
	metav1.Object
}

// Context represents a context of the Reconciler
type Context interface {

	// NewControllerBuilder returns a new builder to create a controller
	NewControllerBuilder() *builder.Builder

	// Client returns the underlying client
	Client() client.Client

	// Scheme returns the underlying scheme
	Scheme() *runtime.Scheme

	// Scheme returns the underlying logger
	Logger() logr.Logger

	// Run checks if the reconciliation can be done and call the reconcile function to do so
	Run(req reconcile.Request, runtimeObject KubeRuntimeObject, reconcile func() error) (reconcile.Result, error)

	// SetOwnershipReference set ownership of the controlled object to the owner
	SetOwnershipReference(owner metav1.Object, controlled metav1.Object) error

	// GetResource is a helper to method to get a resource and do something about its availability
	GetResource(key client.ObjectKey, object runtime.Object, foundCallback func() (err error), notFoundCallback func() (err error)) error
}
