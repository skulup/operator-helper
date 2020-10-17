/*
 * Copyright 2020 Skulup Ltd, Open Collaborators
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
