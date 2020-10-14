package service

import (
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// New2 creates a new service
func New2(namespace, name string, hasClusterIp bool,
	labels map[string]string, servicePorts []v12.ServicePort) *v12.Service {
	if !hasClusterIp {
		labels["headless"] = "true"
	}
	clusterIp := ""
	if !hasClusterIp {
		clusterIp = v12.ClusterIPNone
	}

	return New(namespace, name, labels, v12.ServiceSpec{
		ClusterIP: clusterIp,
		Selector:  labels,
		Ports:     servicePorts,
	})
}

// New creates a new service
func New(namespace, name string, labels map[string]string, spec v12.ServiceSpec) *v12.Service {
	return &v12.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
			Labels:    labels,
		},
		Spec: spec,
	}
}
