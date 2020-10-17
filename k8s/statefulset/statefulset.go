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

package statefulset

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewSpec(replicas int32, service string, selectorMatchLabel map[string]string,
	pvcs []v12.PersistentVolumeClaim, podTemplateSpec v12.PodTemplateSpec) v1.StatefulSetSpec {
	return v1.StatefulSetSpec{
		Replicas: &replicas,
		Selector: &metav1.LabelSelector{
			MatchLabels:      selectorMatchLabel,
			MatchExpressions: nil,
		},
		VolumeClaimTemplates: pvcs,
		ServiceName:          service,
		Template:             podTemplateSpec,
		PodManagementPolicy:  v1.OrderedReadyPodManagement,
		UpdateStrategy: v1.StatefulSetUpdateStrategy{
			Type: v1.RollingUpdateStatefulSetStrategyType,
		},
	}
}

// New creates a new statefulset
func New(namespace, name string, labels map[string]string, spec v1.StatefulSetSpec) *v1.StatefulSet {
	return &v1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
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

// IsReady checks if a statefulset is ready by comparing the desired replicas to the ready replicas
func IsReady(client client.Client, namespace, name string, replicas int32) bool {
	sset := &v1.StatefulSet{}
	err := client.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, sset)
	if err != nil && !errors.IsNotFound(err) {
		fmt.Println(fmt.Sprintf("There was an error on probing for the statefulset: %s", err))
	}
	return err == nil && replicas == sset.Status.ReadyReplicas
}
