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

package pod

import (
	"github.com/skulup/operator-helper/basetype"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewSpec(cfg basetype.PodConfig, volumes []v1.Volume, initContainers []v1.Container, containers []v1.Container) v1.PodSpec {
	var activeDeadlineSeconds *int64
	if cfg.ActiveDeadlineSeconds > 0 {
		activeDeadlineSeconds = &cfg.ActiveDeadlineSeconds
	}
	return v1.PodSpec{
		Affinity:              &cfg.Affinity,
		Tolerations:           cfg.Tolerations,
		NodeSelector:          cfg.NodeSelector,
		RestartPolicy:         cfg.RestartPolicy,
		SecurityContext:       &cfg.SecurityContext,
		ActiveDeadlineSeconds: activeDeadlineSeconds,
		InitContainers:        initContainers,
		Containers:            containers,
		Volumes:               volumes,
	}
}

func NewTemplateSpec(name, generateName string, labels, annotations map[string]string, podSpec v1.PodSpec) v1.PodTemplateSpec {
	return v1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Name:         name,
			GenerateName: generateName,
			Labels:       labels,
			Annotations:  annotations,
		},
		Spec: podSpec,
	}
}
