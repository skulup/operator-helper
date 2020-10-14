package pvcs

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// New creates a new pvc
func New(namespace, name string, labels map[string]string, spec v1.PersistentVolumeClaimSpec) v1.PersistentVolumeClaim {
	return v1.PersistentVolumeClaim{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
			Labels:    labels,
		},
		Spec:   spec,
		Status: v1.PersistentVolumeClaimStatus{},
	}
}
