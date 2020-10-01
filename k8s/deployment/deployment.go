package deployment

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// New creates a new deployment
func New(namespace, name string, labels map[string]string, spec v1.DeploymentSpec) *v1.Deployment {
	return &v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
			Labels:    labels,
		},
		Spec: spec,
	}
}

// IsReady checks if a deployment is ready by comparing the desired replicas to the ready replicas
func IsReady(client client.Client, namespace, name string, replicas int32) bool {
	dep := &v1.Deployment{}
	err := client.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, dep)
	if err != nil && !errors.IsNotFound(err) {
		fmt.Println(fmt.Sprintf("There was an error on probing for the deployment: %s", err))
	}
	return err == nil && replicas == dep.Status.ReadyReplicas
}
